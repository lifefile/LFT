package state

import (
	"github.com/lifefile/LFT/muxdb"
	"github.com/lifefile/LFT/thor"
)

// Stater is the state creator.
type Stater struct {
	db *muxdb.MuxDB
}

// NewStater create a new stater.
func NewStater(db *muxdb.MuxDB) *Stater {
	return &Stater{db}
}

// NewState create a new state object.
func (s *Stater) NewState(root thor.Bytes32) *State {
	return New(s.db, root)
}
