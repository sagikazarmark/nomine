package main // import "github.com/sagikazarmark/nomine"

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"github.com/sagikazarmark/healthz"
	"github.com/sagikazarmark/nomine/api"
	"github.com/sagikazarmark/nomine/app"
	"github.com/sagikazarmark/nomine/services"
	"github.com/sagikazarmark/serverz"
)

func main() {
	defer shutdown.Handle()

	flag.Parse()

	logger.WithFields(logrus.Fields{
		"version":     app.Version,
		"commitHash":  app.CommitHash,
		"buildDate":   app.BuildDate,
		"environment": config.Environment,
	}).Printf("Starting %s service", app.FriendlyServiceName)

	w := logger.Logger.WriterLevel(logrus.ErrorLevel)
	shutdown.Register(w.Close)

	serverManager := serverz.NewServerManager(logger)
	errChan := make(chan error, 10)
	signalChan := make(chan os.Signal, 1)

	var debugServer serverz.Server
	if config.Debug {
		debugServer = &serverz.NamedServer{
			Server: &http.Server{
				Handler:  http.DefaultServeMux,
				ErrorLog: log.New(w, "debug: ", 0),
			},
			Name: "debug",
		}
		shutdown.RegisterAsFirst(debugServer.Close)

		go serverManager.ListenAndStartServer(debugServer, config.DebugAddr)(errChan)
	}

	anaconda.SetConsumerKey(config.TwitterConsumerKey)
	anaconda.SetConsumerSecret(config.TwitterConsumerSecret)
	service := app.NewService(map[string]services.NameChecker{
		"github": services.NewGithub(config.GithubToken, logger),
		"twitter": services.NewTwitter(
			anaconda.NewTwitterApi(config.TwitterAccessKey, config.TwitterAccessSecret),
			logger,
		),
		"docker":     services.NewDocker(logger),
		"com_domain": services.NewWhoisxml("com", config.WhoisxmlUser, config.WhoisxmlPassword, logger),
		"io_domain":  services.NewWhoisxml("io", config.WhoisxmlUser, config.WhoisxmlPassword, logger),
		"org_domain": services.NewWhoisxml("org", config.WhoisxmlUser, config.WhoisxmlPassword, logger),
		"net_domain": services.NewWhoisxml("net", config.WhoisxmlUser, config.WhoisxmlPassword, logger),
	})
	grpcServer := grpc.NewServer()
	api.RegisterNomineServer(grpcServer, service)
	grpcServerWrapper := &serverz.NamedServer{
		Server: &serverz.GrpcServer{grpcServer},
		Name:   "grpc",
	}

	ctx, cancel := context.WithCancel(context.Background())
	restHandler := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterNomineHandlerFromEndpoint(ctx, restHandler, config.GrpcServiceAddr, opts)
	if err != nil {
		logger.Error(err)
	}

	restServer := &serverz.NamedServer{
		Server: &http.Server{
			Handler:  cors.Default().Handler(restHandler),
			ErrorLog: log.New(w, "rest: ", 0),
		},
		Name: "rest",
	}
	shutdown.RegisterAsFirst(serverz.ShutdownFunc(cancel))

	status := healthz.NewStatusChecker(healthz.Healthy)
	readiness := status
	healthHandler := healthz.NewHealthServiceHandler(healthz.NewCheckers(), readiness)
	healthServer := &serverz.NamedServer{
		Server: &http.Server{
			Handler:  healthHandler,
			ErrorLog: log.New(w, "health: ", 0),
		},
		Name: "health",
	}
	shutdown.RegisterAsFirst(healthServer.Close, serverz.ShutdownFunc(grpcServer.Stop), restServer.Close)

	go serverManager.ListenAndStartServer(healthServer, config.HealthAddr)(errChan)
	go serverManager.ListenAndStartServer(grpcServerWrapper, config.GrpcServiceAddr)(errChan)
	go serverManager.ListenAndStartServer(restServer, config.RestServiceAddr)(errChan)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

MainLoop:
	for {
		select {
		case err := <-errChan:
			status.SetStatus(healthz.Unhealthy)

			if err != nil {
				logger.Error(err)
			} else {
				logger.Warning("Error channel received non-error value")
			}

			// Break the loop, proceed with regular shutdown
			break MainLoop
		case s := <-signalChan:
			logger.Infof(fmt.Sprintf("Captured %v", s))
			status.SetStatus(healthz.Unhealthy)

			ctx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
			wg := &sync.WaitGroup{}

			if config.Debug {
				go serverManager.StopServer(debugServer, wg)(ctx)
			}

			go serverManager.StopServer(restServer, wg)(ctx)
			go serverManager.StopServer(grpcServerWrapper, wg)(ctx)
			go serverManager.StopServer(healthServer, wg)(ctx)

			wg.Wait()

			// Cancel context if shutdown completed earlier
			cancel()

			// Break the loop, proceed with regular shutdown
			break MainLoop
		}
	}

	close(errChan)
	close(signalChan)
}
