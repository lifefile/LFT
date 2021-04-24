package energy

import (
	"math/big"
)

type (
	initialSupply struct {
		Token     *big.Int
		Energy    *big.Int
		BlockTime uint64
	}
	totalAddSub struct {
		TotalAdd *big.Int
		TotalSub *big.Int
	}
)
