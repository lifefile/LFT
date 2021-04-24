package state

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/lifefile/LFT/muxdb"
	"github.com/lifefile/LFT/thor"
)

func TestCachedObject(t *testing.T) {
	db := muxdb.NewMem()
	addr := thor.Address{}

	stgTrie := db.NewSecureTrie(StorageTrieName(thor.Blake2b(addr[:])), thor.Bytes32{})
	storages := []struct {
		k thor.Bytes32
		v rlp.RawValue
	}{
		{thor.BytesToBytes32([]byte("key1")), []byte("value1")},
		{thor.BytesToBytes32([]byte("key2")), []byte("value2")},
		{thor.BytesToBytes32([]byte("key3")), []byte("value3")},
		{thor.BytesToBytes32([]byte("key4")), []byte("value4")},
	}

	for _, s := range storages {
		saveStorage(stgTrie, s.k, s.v)
	}

	storageRoot, _ := stgTrie.Commit()

	code := make([]byte, 100)
	rand.Read(code)

	codeHash := crypto.Keccak256(code)
	db.NewStore(codeStoreName).Put(codeHash, code)

	account := Account{
		Balance:     &big.Int{},
		CodeHash:    codeHash,
		StorageRoot: storageRoot[:],
	}

	obj := newCachedObject(db, addr, &account)

	assert.Equal(t,
		M(code, nil),
		M(obj.GetCode()))

	for _, s := range storages {
		assert.Equal(t,
			M(s.v, nil),
			M(obj.GetStorage(s.k)))

	}
}
