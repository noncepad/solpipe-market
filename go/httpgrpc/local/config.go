package client

import "github.com/noncepad/solpipe-market/go/httpgrpc"

type Configuration struct {
	Local *httpgrpc.Address `json:"local"`
}
