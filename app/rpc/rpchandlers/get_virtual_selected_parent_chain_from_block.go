package rpchandlers

import (
	"github.com/kasv2/kasv2d/app/appmessage"
	"github.com/kasv2/kasv2d/app/rpc/rpccontext"
	"github.com/kasv2/kasv2d/domain/consensus/model/externalapi"
	"github.com/kasv2/kasv2d/infrastructure/network/netadapter/router"
)

// HandleGetVirtualSelectedParentChainFromBlock handles the respectively named RPC command
func HandleGetVirtualSelectedParentChainFromBlock(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	getVirtualSelectedParentChainFromBlockRequest := request.(*appmessage.GetVirtualSelectedParentChainFromBlockRequestMessage)

	startHash, err := externalapi.NewDomainHashFromString(getVirtualSelectedParentChainFromBlockRequest.StartHash)
	if err != nil {
		errorMessage := &appmessage.GetVirtualSelectedParentChainFromBlockResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Could not parse startHash: %s", err)
		return errorMessage, nil
	}

	virtualSelectedParentChain, err := context.Domain.Consensus().GetVirtualSelectedParentChainFromBlock(startHash)
	if err != nil {
		response := &appmessage.GetVirtualSelectedParentChainFromBlockResponseMessage{}
		response.Error = appmessage.RPCErrorf("Could not build virtual "+
			"selected parent chain from %s: %s", getVirtualSelectedParentChainFromBlockRequest.StartHash, err)
		return response, nil
	}

	chainChangedNotification, err := context.ConvertVirtualSelectedParentChainChangesToChainChangedNotificationMessage(
		virtualSelectedParentChain, getVirtualSelectedParentChainFromBlockRequest.IncludeAcceptedTransactionIDs)
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetVirtualSelectedParentChainFromBlockResponseMessage(
		chainChangedNotification.RemovedChainBlockHashes, chainChangedNotification.AddedChainBlockHashes,
		chainChangedNotification.AcceptedTransactionIDs)
	return response, nil
}
