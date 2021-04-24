package poa

import (
	"github.com/lifefile/LFT/thor"
)

// Proposer address with status.
type Proposer struct {
	Address thor.Address
	Active  bool
}
