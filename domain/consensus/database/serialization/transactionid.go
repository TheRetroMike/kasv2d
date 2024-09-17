package serialization

import (
	"github.com/kasv2/kasv2d/domain/consensus/model/externalapi"
	"github.com/kasv2/kasv2d/domain/consensus/utils/transactionid"
)

// DbTransactionIDToDomainTransactionID converts DbTransactionId to DomainTransactionID
func DbTransactionIDToDomainTransactionID(dbTransactionID *DbTransactionId) (*externalapi.DomainTransactionID, error) {
	return transactionid.FromBytes(dbTransactionID.TransactionId)
}

// DomainTransactionIDToDbTransactionID converts DomainTransactionID to DbTransactionId
func DomainTransactionIDToDbTransactionID(domainTransactionID *externalapi.DomainTransactionID) *DbTransactionId {
	return &DbTransactionId{TransactionId: domainTransactionID.ByteSlice()}
}
