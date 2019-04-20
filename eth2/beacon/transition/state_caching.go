package transition

import (
	. "github.com/protolambda/zrnt/eth2/beacon"
	. "github.com/protolambda/zrnt/eth2/core"
	"github.com/protolambda/zrnt/eth2/util/ssz"
)

func CacheState(state *BeaconState) {
	prevSlotStateRoot := ssz.HashTreeRoot(state)
	// store the previous slot's post state transition root
	state.LatestStateRoots[state.Slot%SLOTS_PER_HISTORICAL_ROOT] = prevSlotStateRoot

	// cache state root in stored latest_block_header if empty
	if state.LatestBlockHeader.StateRoot == (Root{}) {
		state.LatestBlockHeader.StateRoot = prevSlotStateRoot
	}

	// store latest known block for previous slot
	state.LatestBlockRoots[state.Slot%SLOTS_PER_HISTORICAL_ROOT] = ssz.SigningRoot(state.LatestBlockHeader)
}
