package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	api "github.com/ssdbkey/grpc-challenge/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BlockInfo struct {
	Height string `json:"height,omitempty"`
	Hash   string `json:"hash,omitempty"`
}

type TestResult struct {
	TestResult []BlockInfo `json:"test_result,omitempty"`
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewTendermintClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	var result TestResult

	r, err := c.GetLatestBlock(ctx, &api.GetLatestBlockRequest{})
	if err != nil {
		log.Fatalf("could not get block: %v", err)
	}

	result.TestResult = append(result.TestResult, BlockInfo{
		Height: r.Block.Header.Height,
		Hash:   string(r.BlockId.Hash),
	})

	for len(result.TestResult) < 5 {
		r, err := c.GetLatestBlock(ctx, &api.GetLatestBlockRequest{})
		if err != nil {
			log.Fatalf("could not get block: %v", err)
		}

		if result.TestResult[len(result.TestResult)-1].Height != r.Block.Header.Height {
			result.TestResult = append(result.TestResult, BlockInfo{
				Height: r.Block.Header.Height,
				Hash:   string(r.BlockId.Hash),
			})
		}

		time.Sleep(1 * time.Second)
	}

	file, _ := json.MarshalIndent(result, "", " ")
	_ = ioutil.WriteFile("exported.json", file, 0644)
}
