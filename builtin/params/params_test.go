package params

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/lifefile/LFT/muxdb"
	"github.com/lifefile/LFT/state"
	"github.com/lifefile/LFT/thor"
)

func TestParamsGetSet(t *testing.T) {
	db := muxdb.NewMem()
	st := state.New(db, thor.Bytes32{})
	setv := big.NewInt(10)
	key := thor.BytesToBytes32([]byte("key"))
	p := New(thor.BytesToAddress([]byte("par")), st)
	p.Set(key, setv)

	getv, err := p.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, setv, getv)
}
