package testapi

import (
	"github.com/kasv2/kasv2d/domain/consensus/model"
	"github.com/kasv2/kasv2d/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
