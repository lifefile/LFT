package tx

import (
	"math/big"

	"github.com/lifefile/LFT/thor"
)

// Transfer token transfer log.
type Transfer struct {
	Sender    thor.Address
	Recipient thor.Address
	Amount    *big.Int
}

// Transfers slisce of transfer logs.
type Transfers []*Transfer
