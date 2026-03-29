package vecmt

import (
	"fmt"
	"sort"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"

	"github.com/Fantom-foundation/go-opera/inter"
)

// medianTimeIndex is a handy index for the MedianTime() func
type medianTimeIndex struct {
	weight       pos.Weight
	creationTime inter.Timestamp
}

// MedianTime calculates weighted median of claimed time within highest observed events.
func (vi *Index) MedianTime(id hash.Event, defaultTime inter.Timestamp) inter.Timestamp {
	vi.Engine.InitBranchesInfo()
	// Get event by hash
	_before := vi.Engine.GetMergedHighestBefore(id)
	if _before == nil {
		vi.crit(fmt.Errorf("event=%s not found", id.String()))
	}
	before := _before.(*HighestBefore)

	// honestTotalWeight sums only non-cheater validators. Bounded by
	// validators.TotalWeight() which is capped at construction time by the
	// SFC's validator set size and max-stake limits, so overflow is not
	// possible under correct validator set construction.
	honestTotalWeight := pos.Weight(0)
	highests := make([]medianTimeIndex, 0, len(vi.validatorIdxs))
	// convert []HighestBefore -> []medianTimeIndex
	for creatorIdxI := range vi.validators.IDs() {
		creatorIdx := idx.Validator(creatorIdxI)
		highest := medianTimeIndex{}
		highest.weight = vi.validators.GetWeightByIdx(creatorIdx)
		highest.creationTime = before.VTime.Get(creatorIdx)
		seq := before.VSeq.Get(creatorIdx)

		// edge cases
		if seq.IsForkDetected() {
			// cheaters don't influence medianTime
			highest.weight = 0
		} else if seq.Seq == 0 {
			// if no event was observed from this node, then use genesisTime
			highest.creationTime = defaultTime
		}

		highests = append(highests, highest)
		honestTotalWeight += highest.weight
	}
	// If all validators are detected as cheaters, no honest weight exists.
	// Return the default time as a safe fallback in this catastrophic scenario.
	if honestTotalWeight == 0 {
		return defaultTime
	}

	// sort by claimed time
	if vi.elemont {
		// Post-Elemont: stable sort with weight tie-breaker for determinism
		sort.SliceStable(highests, func(i, j int) bool {
			a, b := highests[i], highests[j]
			if a.creationTime != b.creationTime {
				return a.creationTime < b.creationTime
			}
			return a.weight < b.weight
		})
	} else {
		// Pre-Elemont: unstable sort by time only (legacy behavior)
		sort.Slice(highests, func(i, j int) bool {
			return highests[i].creationTime < highests[j].creationTime
		})
	}

	// Calculate weighted median
	halfWeight := honestTotalWeight / 2
	if halfWeight == 0 {
		halfWeight = 1
	}
	var currWeight pos.Weight
	var median inter.Timestamp
	for _, highest := range highests {
		currWeight += highest.weight
		if currWeight >= halfWeight {
			median = highest.creationTime
			break
		}
	}

	// sanity check
	if currWeight < halfWeight || currWeight > honestTotalWeight {
		vi.crit(fmt.Errorf("median wasn't calculated correctly, median=%d, currWeight=%d, totalWeight=%d, len(highests)=%d, id=%s",
			median,
			currWeight,
			honestTotalWeight,
			len(highests),
			id.String(),
		))
	}

	return median
}
