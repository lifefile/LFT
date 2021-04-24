package tx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/lifefile/LFT/thor"
	"github.com/lifefile/LFT/trie"
)

// Receipt represents the results of a transaction.
type Receipt struct {
	// gas used by this tx
	GasUsed uint64
	// the one who paid for gas
	GasPayer thor.Address
	// energy paid for used gas
	Paid *big.Int
	// energy reward given to block proposer
	Reward *big.Int
	// if the tx reverted
	Reverted bool
	// outputs of clauses in tx
	Outputs []*Output
}

// Output output of clause execution.
type Output struct {
	// events produced by the clause
	Events Events
	// transfer occurred in clause
	Transfers Transfers
}

// Receipts slice of receipts.
type Receipts []*Receipt

// RootHash computes merkle root hash of receipts.
func (rs Receipts) RootHash() thor.Bytes32 {
	if len(rs) == 0 {
		// optimized
		return emptyRoot
	}
	return trie.DeriveRoot(derivableReceipts(rs))
}

// implements DerivableList
type derivableReceipts Receipts

func (rs derivableReceipts) Len() int {
	return len(rs)
}
func (rs derivableReceipts) GetRlp(i int) []byte {
	data, err := rlp.EncodeToBytes(rs[i])
	if err != nil {
		panic(err)
	}
	return data
}
