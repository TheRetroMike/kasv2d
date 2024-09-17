package rpchandlers

import (
	"github.com/kasv2/kasv2d/app/appmessage"
	"github.com/kasv2/kasv2d/app/rpc/rpccontext"
	"github.com/kasv2/kasv2d/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
