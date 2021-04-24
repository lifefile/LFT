package node

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/mclock"
	"github.com/lifefile/LFT/block"
	"github.com/lifefile/LFT/thor"
)

type blockStats struct {
	exec, commit               mclock.AbsTime
	txs                        int
	usedGas                    uint64
	processed, queued, ignored int
}

func (s *blockStats) UpdateProcessed(n int, txs int, exec, commit mclock.AbsTime, usedGas uint64) {
	s.processed += n
	s.txs += txs
	s.exec += exec
	s.commit += commit
	s.usedGas += usedGas
}

func (s *blockStats) UpdateIgnored(n int) {
	s.ignored += n
}

func (s *blockStats) UpdateQueued(n int) {
	s.queued += n
}

func (s *blockStats) LogContext(last *block.Header) []interface{} {
	return []interface{}{
		"txs", s.txs,
		"mgas", float64(s.usedGas) / 1000 / 1000,
		"et", fmt.Sprintf("%v|%v", common.PrettyDuration(s.exec), common.PrettyDuration(s.commit)),
		"mgas/s", float64(s.usedGas) * 1000 / float64(s.exec+s.commit),
		"id", shortID(last.ID()),
	}
}

func shortID(id thor.Bytes32) string {
	return fmt.Sprintf("[#%v…%x]", block.Number(id), id[28:])
}
