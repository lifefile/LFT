package txpool

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/lifefile/LFT/block"
	"github.com/lifefile/LFT/chain"
	"github.com/lifefile/LFT/genesis"
	"github.com/lifefile/LFT/muxdb"
	"github.com/lifefile/LFT/state"
	"github.com/lifefile/LFT/thor"
	"github.com/lifefile/LFT/tx"
)

func newChainRepo(db *muxdb.MuxDB) *chain.Repository {
	gene := genesis.NewDevnet()
	b0, _, _, _ := gene.Build(state.NewStater(db))
	repo, _ := chain.NewRepository(db, b0)
	return repo
}

func signTx(tx *tx.Transaction, acc genesis.DevAccount) *tx.Transaction {
	sig, _ := crypto.Sign(tx.SigningHash().Bytes(), acc.PrivateKey)
	return tx.WithSignature(sig)
}

func newTx(chainTag byte, clauses []*tx.Clause, gas uint64, blockRef tx.BlockRef, expiration uint32, dependsOn *thor.Bytes32, features tx.Features, from genesis.DevAccount) *tx.Transaction {
	builder := new(tx.Builder).ChainTag(chainTag)
	for _, c := range clauses {
		builder.Clause(c)
	}

	tx := builder.BlockRef(blockRef).
		Expiration(expiration).
		Nonce(rand.Uint64()).
		DependsOn(dependsOn).
		Features(features).
		Gas(gas).Build()

	return signTx(tx, from)
}

func TestSort(t *testing.T) {
	objs := []*txObject{
		{overallGasPrice: big.NewInt(10)},
		{overallGasPrice: big.NewInt(20)},
		{overallGasPrice: big.NewInt(30)},
	}
	sortTxObjsByOverallGasPriceDesc(objs)

	assert.Equal(t, big.NewInt(30), objs[0].overallGasPrice)
	assert.Equal(t, big.NewInt(20), objs[1].overallGasPrice)
	assert.Equal(t, big.NewInt(10), objs[2].overallGasPrice)
}

func TestResolve(t *testing.T) {
	acc := genesis.DevAccounts()[0]
	tx := newTx(0, nil, 21000, tx.BlockRef{}, 100, nil, tx.Features(0), acc)

	txObj, err := resolveTx(tx, false)
	assert.Nil(t, err)
	assert.Equal(t, tx, txObj.Transaction)

	assert.Equal(t, acc.Address, txObj.Origin())

}

func TestExecutable(t *testing.T) {
	acc := genesis.DevAccounts()[0]

	db := muxdb.NewMem()
	repo := newChainRepo(db)
	b0 := repo.GenesisBlock()
	b1 := new(block.Builder).ParentID(b0.Header().ID()).GasLimit(10000000).TotalScore(100).Build()
	repo.AddBlock(b1, nil)
	st := state.New(db, repo.GenesisBlock().Header().StateRoot())

	tests := []struct {
		tx          *tx.Transaction
		expected    bool
		expectedErr string
	}{
		{newTx(0, nil, 21000, tx.BlockRef{}, 100, nil, tx.Features(0), acc), true, ""},
		{newTx(0, nil, math.MaxUint64, tx.BlockRef{}, 100, nil, tx.Features(0), acc), false, "gas too large"},
		{newTx(0, nil, 21000, tx.BlockRef{1}, 100, nil, tx.Features(0), acc), true, "block ref out of schedule"},
		{newTx(0, nil, 21000, tx.BlockRef{0}, 0, nil, tx.Features(0), acc), true, "expired"},
		{newTx(0, nil, 21000, tx.BlockRef{0}, 100, &thor.Bytes32{}, tx.Features(0), acc), false, ""},
	}

	for _, tt := range tests {
		txObj, err := resolveTx(tt.tx, false)
		assert.Nil(t, err)

		exe, err := txObj.Executable(repo.NewChain(b1.Header().ID()), st, b1.Header())
		if tt.expectedErr != "" {
			assert.Equal(t, tt.expectedErr, err.Error())
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, exe)
		}
	}
}