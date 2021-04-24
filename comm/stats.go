package comm

import (
	"github.com/lifefile/LFT/thor"
)

// type Traffic struct {
// 	Bytes    uint64
// 	Requests uint64
// 	Errors   uint64
// }

// PeerStats records stats of a peer.
type PeerStats struct {
	Name        string
	BestBlockID thor.Bytes32
	TotalScore  uint64
	PeerID      string
	NetAddr     string
	Inbound     bool
	Duration    uint64 // in seconds
}
