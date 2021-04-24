package state

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/lifefile/LFT/muxdb"
	"github.com/lifefile/LFT/thor"
)

func TestStage(t *testing.T) {
	db := muxdb.NewMem()

	state := New(db, thor.Bytes32{})

	addr := thor.BytesToAddress([]byte("acc1"))

	balance := big.NewInt(10)
	code := []byte{1, 2, 3}

	storage := map[thor.Bytes32]thor.Bytes32{
		thor.BytesToBytes32([]byte("s1")): thor.BytesToBytes32([]byte("v1")),
		thor.BytesToBytes32([]byte("s2")): thor.BytesToBytes32([]byte("v2")),
		thor.BytesToBytes32([]byte("s3")): thor.BytesToBytes32([]byte("v3"))}

	state.SetBalance(addr, balance)
	state.SetCode(addr, code)
	for k, v := range storage {
		state.SetStorage(addr, k, v)
	}

	stage, err := state.Stage()
	assert.Nil(t, err)

	hash := stage.Hash()

	root, err := stage.Commit()
	assert.Nil(t, err)

	assert.Equal(t, hash, root)

	state = New(db, root)

	assert.Equal(t, M(balance, nil), M(state.GetBalance(addr)))
	assert.Equal(t, M(code, nil), M(state.GetCode(addr)))
	assert.Equal(t, M(thor.Bytes32(crypto.Keccak256Hash(code)), nil), M(state.GetCodeHash(addr)))

	for k, v := range storage {
		assert.Equal(t, M(v, nil), M(state.GetStorage(addr, k)))
	}
}
