package tendermint

import (
	"context"

	"github.com/ssdbkey/grpc-challenge/pkg/api"
)

// GRPCServer struct
type GRPCServer struct{}

func (s *GRPCServer) Health(ctx context.Context, req *api.EmptyRequest) (*api.AddResponse, error) {
	return &api.AddResponse{
		Result: req.GetX() + req.GetY(),
	}, nil
}
