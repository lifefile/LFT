package state

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/lifefile/LFT/muxdb"
	"github.com/lifefile/LFT/thor"
)

func M(a ...interface{}) []interface{} {
	return a
}

func TestAccount(t *testing.T) {
	assert.True(t, emptyAccount().IsEmpty())

	acc := emptyAccount()
	acc.Balance = big.NewInt(1)
	assert.False(t, acc.IsEmpty())
	acc = emptyAccount()
	acc.CodeHash = []byte{1}
	assert.False(t, acc.IsEmpty())

	acc = emptyAccount()
	acc.Energy = big.NewInt(1)
	assert.False(t, acc.IsEmpty())

	acc = emptyAccount()
	acc.StorageRoot = []byte{1}
	assert.True(t, acc.IsEmpty())
}

func newTrie() *muxdb.Trie {
	return muxdb.NewMem().NewSecureTrie("", thor.Bytes32{})
}
func TestTrie(t *testing.T) {
	trie := newTrie()

	addr := thor.BytesToAddress([]byte("account1"))
	assert.Equal(t,
		M(loadAccount(trie, addr)),
		M(emptyAccount(), nil),
		"should load an empty account")

	acc1 := Account{
		big.NewInt(1),
		big.NewInt(0),
		0,
		[]byte("master"),
		[]byte("code hash"),
		[]byte("storage root"),
	}
	saveAccount(trie, addr, &acc1)
	assert.Equal(t,
		M(loadAccount(trie, addr)),
		M(&acc1, nil))

	saveAccount(trie, addr, emptyAccount())
	assert.Equal(t,
		M(trie.Get(addr[:])),
		M([]byte(nil), nil),
		"empty account should be deleted")
}

func TestStorageTrie(t *testing.T) {
	trie := newTrie()

	key := thor.BytesToBytes32([]byte("key"))
	assert.Equal(t,
		M(loadStorage(trie, key)),
		M(rlp.RawValue(nil), nil))

	value := rlp.RawValue("value")
	saveStorage(trie, key, value)
	assert.Equal(t,
		M(loadStorage(trie, key)),
		M(value, nil))

	saveStorage(trie, key, nil)
	assert.Equal(t,
		M(trie.Get(key[:])),
		M([]byte(nil), nil),
		"empty storage value should be deleted")
}