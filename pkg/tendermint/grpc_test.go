package tendermint

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	api "github.com/ssdbkey/grpc-challenge/pkg/api"
	"github.com/test-go/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TendermintTestSuite struct {
	suite.Suite
}

func (suite *TendermintTestSuite) SetupTest() {}

func (suite *TendermintTestSuite) TestGetBlockByHeight() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewTendermintClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	r, err := c.GetBlockByHeight(ctx, &api.GetBlockByHeightRequest{Height: 100})
	if err != nil {
		log.Fatalf("could not get block: %v", err)
	}

	// Prepare your HTTP request
	jsonBody := map[string]interface{}{
		"height": 100,
	}
	reqBody, _ := json.Marshal(jsonBody)
	httpReq, err := http.NewRequest(http.MethodGet, "https://rpc.osmosis.zone/block", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatalf("could not get direct client api response: %v", err)
	}

	// Send your HTTP request
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	defer httpResp.Body.Close()

	var resp HttpResponse
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		log.Fatalf("response decode failed: %v", err)
	}

	suite.Require().Equal(&r, &resp.Result)
}

func TestTendermintTestSuite(t *testing.T) {
	testSuite := *new(TendermintTestSuite)

	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &GRPCServer{}

	// Register gRPC server
	api.RegisterTendermintServer(s, srv)

	// Listen on port 50051
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on %v", l.Addr())

	go s.Serve(l)

	time.Sleep(1 * time.Second)

	suite.Run(t, &testSuite)
	s.Stop()
}
