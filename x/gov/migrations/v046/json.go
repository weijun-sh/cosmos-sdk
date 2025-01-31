package v046

import (
	"github.com/weijun-sh/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/weijun-sh/cosmos-sdk/x/gov/types/v1beta2"
)

// MigrateJSON accepts exported v0.43 x/gov genesis state and migrates it to
// v0.46 x/gov genesis state. The migration includes:
//
// - Updating everything to v1beta2.
// - Migrating proposals to be Msg-based.
func MigrateJSON(oldState *v1beta1.GenesisState) (*v1beta2.GenesisState, error) {
	newProps, err := convertToNewProposals(oldState.Proposals)
	if err != nil {
		return nil, err
	}
	newVotes, err := convertToNewVotes(oldState.Votes)
	if err != nil {
		return nil, err
	}

	depParams, votingParms, tallyParams := convertToNewDepParams(oldState.DepositParams), convertToNewVotingParams(oldState.VotingParams), convertToNewTallyParams(oldState.TallyParams)

	return &v1beta2.GenesisState{
		StartingProposalId: oldState.StartingProposalId,
		Deposits:           convertToNewDeposits(oldState.Deposits),
		Votes:              newVotes,
		Proposals:          newProps,
		DepositParams:      &depParams,
		VotingParams:       &votingParms,
		TallyParams:        &tallyParams,
	}, nil
}
