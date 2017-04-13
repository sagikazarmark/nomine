package app

import (
	"github.com/sagikazarmark/nomine/api"
	"github.com/sagikazarmark/nomine/services"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Service implements the RPC server
type Service struct {
	checkers map[string]services.NameChecker
}

// NewService creates a new Service
func NewService(checkers map[string]services.NameChecker) *Service {
	return &Service{checkers}
}

// Check checks a name availability
func (s *Service) Check(ctx context.Context, request *api.NameCheckRequest) (*api.NameCheckResponse, error) {
	if request.GetName() == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "name should not be empty.")
	}

	if len(request.GetServices()) < 1 {
		return nil, grpc.Errorf(codes.InvalidArgument, "service list should not be empty.")
	}

	response := &api.NameCheckResponse{}
	response.Results = make(map[string]api.Result)

	for _, service := range request.GetServices() {
		if checker, ok := s.checkers[service]; ok {
			result, err := checker.Check(request.GetName())
			if err != nil {
				response.Results[service] = api.Result_UNKOWN
			} else if result {
				response.Results[service] = api.Result_AVAILABLE
			} else {
				response.Results[service] = api.Result_UNAVAILABLE
			}
		} else {
			response.Results[service] = api.Result_UNKOWN
		}
	}

	return response, nil
}