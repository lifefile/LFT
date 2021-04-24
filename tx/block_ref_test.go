package tx_test

import (
	"math/rand"
	"testing"

	"github.com/lifefile/LFT/thor"

	"github.com/stretchr/testify/assert"
	"github.com/lifefile/LFT/tx"
)

func TestBlockRef(t *testing.T) {
	assert.Equal(t, uint32(0), tx.BlockRef{}.Number())

	assert.Equal(t, tx.BlockRef{0, 0, 0, 0xff, 0, 0, 0, 0}, tx.NewBlockRef(0xff))

	var bid thor.Bytes32
	rand.Read(bid[:])

	br := tx.NewBlockRefFromID(bid)
	assert.Equal(t, bid[:8], br[:])
}
