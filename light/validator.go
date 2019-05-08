package light

import (
	"github.com/dexon-foundation/dexon/common"
	"github.com/dexon-foundation/dexon/consensus"
	"github.com/dexon-foundation/dexon/log"
)

type HeaderValidator struct {
	chain *LightChain
}

func (v *HeaderValidator) ValidateWitnessData(height uint64, blockHash common.Hash) error {
	b := v.chain.GetHeaderByNumber(height)
	if b == nil {
		log.Error("can not find block %v either pending or confirmed block", height)
		return consensus.ErrWitnessMismatch
	}

	if b.Hash() != blockHash {
		log.Error("invalid witness block %s vs %s",
			b.Hash().String(), blockHash.String())
		return consensus.ErrWitnessMismatch
	}
	return nil
}
