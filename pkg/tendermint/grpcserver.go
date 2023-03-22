package tendermint

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	api "github.com/ssdbkey/grpc-challenge/pkg/api"
)

// GRPCServer struct
type GRPCServer struct{}

type HttpResponse struct {
	Id      int32                `json:"id,omitempty"`
	Jsonrpc string               `json:"jsonrpc,omitempty"`
	Result  api.GetBlockResponse `json:"result,omitempty"`
}

func (s *GRPCServer) GetLatestBlock(ctx context.Context, req *api.GetLatestBlockRequest) (*api.GetBlockResponse, error) {
	// Prepare your HTTP request
	httpReq, err := http.NewRequest(http.MethodGet, "https://rpc.osmosis.zone/block", nil)
	if err != nil {
		return nil, err
	}

	// Send your HTTP request
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	var resp HttpResponse
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

func (s *GRPCServer) GetBlockByHeight(ctx context.Context, req *api.GetBlockByHeightRequest) (*api.GetBlockResponse, error) {
	// Prepare your HTTP request
	jsonBody := map[string]interface{}{
		"height": req.Height,
	}

	reqBody, err := json.Marshal(jsonBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodGet, "https://rpc.osmosis.zone/block", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	// Send your HTTP request
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	var resp HttpResponse
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp.Result, nil
}
