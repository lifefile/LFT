package genesis_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/lifefile/LFT/genesis"
	"github.com/lifefile/LFT/muxdb"
	"github.com/lifefile/LFT/state"
	"github.com/lifefile/LFT/thor"
)

func TestTestnetGenesis(t *testing.T) {
	db := muxdb.NewMem()
	gene := genesis.NewTestnet()

	b0, _, _, err := gene.Build(state.NewStater(db))
	assert.Nil(t, err)

	st := state.New(db, b0.Header().StateRoot())

	v, err := st.Exists(thor.MustParseAddress("0xe59D475Abe695c7f67a8a2321f33A856B0B4c71d"))
	assert.Nil(t, err)
	assert.True(t, v)
}
