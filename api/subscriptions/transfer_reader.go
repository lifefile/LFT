package subscriptions

import (
	"github.com/lifefile/LFT/chain"
	"github.com/lifefile/LFT/thor"
)

type transferReader struct {
	repo        *chain.Repository
	filter      *TransferFilter
	blockReader chain.BlockReader
}

func newTransferReader(repo *chain.Repository, position thor.Bytes32, filter *TransferFilter) *transferReader {
	return &transferReader{
		repo:        repo,
		filter:      filter,
		blockReader: repo.NewBlockReader(position),
	}
}

func (tr *transferReader) Read() ([]interface{}, bool, error) {
	blocks, err := tr.blockReader.Read()
	if err != nil {
		return nil, false, err
	}
	var msgs []interface{}
	for _, block := range blocks {
		receipts, err := tr.repo.GetBlockReceipts(block.Header().ID())
		if err != nil {
			return nil, false, err
		}
		txs := block.Transactions()
		for i, receipt := range receipts {
			for j, output := range receipt.Outputs {
				for _, transfer := range output.Transfers {
					origin, err := txs[i].Origin()
					if err != nil {
						return nil, false, err
					}
					if tr.filter.Match(transfer, origin) {
						msg, err := convertTransfer(block.Header(), txs[i], uint32(j), transfer, block.Obsolete)
						if err != nil {
							return nil, false, err
						}
						msgs = append(msgs, msg)
					}
				}
			}
		}
	}
	return msgs, len(blocks) > 0, nil
}
