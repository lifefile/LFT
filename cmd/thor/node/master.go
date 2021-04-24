package node

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lifefile/LFT/thor"
)

type Master struct {
	PrivateKey  *ecdsa.PrivateKey
	Beneficiary *thor.Address
}

func (m *Master) Address() thor.Address {
	return thor.Address(crypto.PubkeyToAddress(m.PrivateKey.PublicKey))
}
