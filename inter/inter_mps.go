package inter

import (
	"errors"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

const (
	// MinAccomplicesForProof defines how many validators must have signed the same wrong vote.
	// Otherwise, wrong-signer is not liable as a protection against singular software/hardware failures
	MinAccomplicesForProof = 2
)

type EventsDoublesign struct {
	Pair [2]SignedEventLocator
}

type BlockVoteDoublesign struct {
	Block idx.Block
	Pair  [2]LlrSignedBlockVotes
}

func (p BlockVoteDoublesign) GetVote(i int) hash.Hash {
	if p.Block < p.Pair[i].Val.Start {
		return hash.Hash{}
	}
	off := p.Block - p.Pair[i].Val.Start
	if off >= idx.Block(len(p.Pair[i].Val.Votes)) {
		return hash.Hash{}
	}
	return p.Pair[i].Val.Votes[off]
}

type WrongBlockVote struct {
	Block      idx.Block
	Pals       [MinAccomplicesForProof]LlrSignedBlockVotes
	WrongEpoch bool
}

func (p WrongBlockVote) GetVote(i int) hash.Hash {
	if p.Block < p.Pals[i].Val.Start {
		return hash.Hash{}
	}
	off := p.Block - p.Pals[i].Val.Start
	if off >= idx.Block(len(p.Pals[i].Val.Votes)) {
		return hash.Hash{}
	}
	return p.Pals[i].Val.Votes[off]
}

type EpochVoteDoublesign struct {
	Pair [2]LlrSignedEpochVote
}

type WrongEpochVote struct {
	Pals [MinAccomplicesForProof]LlrSignedEpochVote
}

const maxMPVotesPerBV = 128 // cap on Votes slice length inside misbehaviour proofs (2x MaxBlockVotesPerEvent)

func validateMPVoteSizes(mps []MisbehaviourProof) error {
	for i := range mps {
		mp := &mps[i]
		if mp.BlockVoteDoublesign != nil {
			for j := range mp.BlockVoteDoublesign.Pair {
				if len(mp.BlockVoteDoublesign.Pair[j].Val.Votes) > maxMPVotesPerBV {
					return errors.New("misbehaviour proof has too many block votes")
				}
			}
		}
		if mp.WrongBlockVote != nil {
			for j := range mp.WrongBlockVote.Pals {
				if len(mp.WrongBlockVote.Pals[j].Val.Votes) > maxMPVotesPerBV {
					return errors.New("misbehaviour proof has too many block votes")
				}
			}
		}
	}
	return nil
}

type MisbehaviourProof struct {
	EventsDoublesign *EventsDoublesign `rlp:"nil"`

	BlockVoteDoublesign *BlockVoteDoublesign `rlp:"nil"`

	WrongBlockVote *WrongBlockVote `rlp:"nil"`

	EpochVoteDoublesign *EpochVoteDoublesign `rlp:"nil"`

	WrongEpochVote *WrongEpochVote `rlp:"nil"`
}
