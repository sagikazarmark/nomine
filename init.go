package main

import (
	"flag"
	"time"

	_ "expvar"
	"net/http"
	_ "net/http/pprof"

	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_fluent"
	"github.com/kelseyhightower/envconfig"
	"github.com/sagikazarmark/nomine/app"
	"github.com/sagikazarmark/serverz"
	"golang.org/x/net/trace"
	"google.golang.org/grpc/grpclog"
)

// Global context variables
var (
	config   = &app.Configuration{}
	logger   = logrus.New().WithField("service", app.ServiceName) // Use logrus.FieldLogger type
	shutdown = serverz.NewShutdown(logger)
)

func init() {
	// Register shutdown handler in logrus
	logrus.RegisterExitHandler(shutdown.Handle)

	// Set global gRPC logger
	grpclog.SetLogger(logger.WithField("server", "grpc"))

	// Load configuration from environment
	err := envconfig.Process("", config)
	if err != nil {
		logger.Fatal(err)
	}

	defaultAddr := ""

	// Listen on loopback interface in development mode
	if config.Environment == "development" {
		defaultAddr = "127.0.0.1"
	}

	// Load flags into configuration
	flag.StringVar(&config.GrpcServiceAddr, "grpc", defaultAddr+":80", "gRPC service address.")
	flag.StringVar(&config.RestServiceAddr, "rest", defaultAddr+":81", "REST service address.")
	flag.StringVar(&config.HealthAddr, "health", defaultAddr+":10000", "Health service address.")
	flag.StringVar(&config.DebugAddr, "debug", defaultAddr+":10001", "Debug service address.")
	flag.DurationVar(&config.ShutdownTimeout, "shutdown", 2*time.Second, "Shutdown timeout.")

	// This is probably OK as the service runs in Docker
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	// Initialize Fluentd
	if config.FluentdEnabled {
		fluentdHook, err := logrus_fluent.New(config.FluentdHost, config.FluentdPort)
		if err != nil {
			logger.Panic(err)
		}

		fluentdHook.SetTag(app.ServiceName)
		fluentdHook.AddFilter("error", logrus_fluent.FilterError)

		logger.Logger.Hooks.Add(fluentdHook)
		shutdown.Register(fluentdHook.Fluent.Close)
	}
}
