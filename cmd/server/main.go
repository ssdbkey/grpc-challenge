package main

import (
	"log"
	"net"

	api "github.com/ssdbkey/grpc-challenge/pkg/api"
	"github.com/ssdbkey/grpc-challenge/pkg/tendermint"
	"google.golang.org/grpc"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &tendermint.GRPCServer{}

	// Register gRPC server
	api.RegisterTendermintServer(s, srv)

	// Listen on port 50051
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on %v", l.Addr())

	// Start gRPC server
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
