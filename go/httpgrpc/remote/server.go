package remote

import (
	"context"

	pbj "github.com/noncepad/solpipe-market/go/proto/solanajsonrpc"
	"google.golang.org/grpc"
)

type external struct {
	pbj.UnimplementedJsonRpcServer
	ctx    context.Context
	config *Configuration
}

func Add(ctx context.Context, s *grpc.Server, config *Configuration) {
	e1 := external{
		ctx:    ctx,
		config: config,
	}
	pbj.RegisterJsonRpcServer(s, e1)
}
