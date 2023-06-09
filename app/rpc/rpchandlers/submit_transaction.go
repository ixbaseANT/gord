package rpchandlers

import (
	"github.com/ixbaseANT/gord/app/appmessage"
	"github.com/ixbaseANT/gord/app/rpc/rpccontext"
	"github.com/ixbaseANT/gord/domain/consensus/utils/consensushashing"
	"github.com/ixbaseANT/gord/domain/miningmanager/mempool"
	"github.com/ixbaseANT/gord/infrastructure/network/netadapter/router"
	"github.com/pkg/errors"
	"fmt"
)

// HandleSubmitTransaction handles the respectively named RPC command
func HandleSubmitTransaction(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	submitTransactionRequest := request.(*appmessage.SubmitTransactionRequestMessage)

	domainTransaction, err := appmessage.RPCTransactionToDomainTransaction(submitTransactionRequest.Transaction)
	if err != nil {
		errorMessage := &appmessage.SubmitTransactionResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Could not parse transaction: %s", err)
		return errorMessage, nil
	}

	transactionID := consensushashing.TransactionID(domainTransaction)
	err = context.ProtocolManager.AddTransaction(domainTransaction, submitTransactionRequest.AllowOrphan)
	if err != nil {
		if !errors.As(err, &mempool.RuleError{}) {
			return nil, err
		}
		log.Debugf("2.Rejected transaction %s: %s", transactionID, err)
		errorMessage := &appmessage.SubmitTransactionResponseMessage{}
fmt.Println("=errorMessage=",errorMessage)
		errorMessage.Error = appmessage.RPCErrorf("-3.Rejected transaction %s: %s", transactionID, err)
		return errorMessage, nil
	}

	response := appmessage.NewSubmitTransactionResponseMessage(transactionID.String())
	return response, nil
}
