package adder

import (
	"context"

	"github.com/ssdbkey/grpc-challenge/pkg/api"
)

// GRPCServer struct
type GRPCServer struct{}

// Add method for calculate X + Y
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{
		Result: req.GetX() + req.GetY(),
	}, nil
}
