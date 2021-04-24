package subscriptions

import (
	"github.com/lifefile/LFT/chain"
	"github.com/lifefile/LFT/thor"
)

type blockReader struct {
	repo        *chain.Repository
	blockReader chain.BlockReader
}

func newBlockReader(repo *chain.Repository, position thor.Bytes32) *blockReader {
	return &blockReader{
		repo:        repo,
		blockReader: repo.NewBlockReader(position),
	}
}

func (br *blockReader) Read() ([]interface{}, bool, error) {
	blocks, err := br.blockReader.Read()
	if err != nil {
		return nil, false, err
	}
	var msgs []interface{}
	for _, block := range blocks {
		msg, err := convertBlock(block)
		if err != nil {
			return nil, false, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, len(blocks) > 0, nil
}
