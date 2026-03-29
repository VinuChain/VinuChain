package vecmt

import (
	"errors"

	"github.com/Fantom-foundation/lachesis-base/hash"
)

// NoCheaters excludes events which are observed by selfParents as cheaters.
// Called by emitter to exclude cheater's events from potential parents list.
func (vi *Index) NoCheaters(selfParent *hash.Event, options hash.Events) hash.Events {
	if selfParent == nil {
		return options
	}
	vi.InitBranchesInfo()

	if !vi.Engine.AtLeastOneFork() {
		return options
	}

	filtered := make(hash.Events, 0, len(options))
	if vi.elemont {
		// Post-Elemont: use merged view so branch indices are collapsed to
		// validator indices. The raw GetHighestBefore is branch-indexed;
		// when forks exist, branch indices don't match validator indices,
		// causing misidentification.
		merged := vi.GetMergedHighestBefore(*selfParent)
		for _, id := range options {
			e := vi.getEvent(id)
			if e == nil {
				vi.crit(errors.New("event not found"))
			}
			if !merged.VSeq.Get(vi.validatorIdxs[e.Creator()]).IsForkDetected() {
				filtered.Add(id)
			}
		}
	} else {
		// Pre-Elemont: use raw branch-indexed HighestBeforeSeq (legacy behavior).
		seqBefore := vi.Base.GetHighestBefore(*selfParent)
		for _, id := range options {
			e := vi.getEvent(id)
			if e == nil {
				vi.crit(errors.New("event not found"))
			}
			creatorIdx := vi.validatorIdxs[e.Creator()]
			if !seqBefore.Get(creatorIdx).IsForkDetected() {
				filtered.Add(id)
			}
		}
	}
	return filtered
}
