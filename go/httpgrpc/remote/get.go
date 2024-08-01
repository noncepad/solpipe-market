package remote

import pbj "github.com/noncepad/solpipe-market/go/proto/solanajsonrpc"

func (e1 external) Get(req *pbj.Header, stream pbj.JsonRpc_GetServer) error {
	return nil
}
