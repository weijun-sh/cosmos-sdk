package v1beta2

import (
	sdk "github.com/weijun-sh/cosmos-sdk/types"
)

// ValidatorGovInfo used for tallying
type ValidatorGovInfo struct {
	Address             sdk.ValAddress      // address of the validator operator
	BondedTokens        sdk.Int             // Power of a Validator
	DelegatorShares     sdk.Dec             // Total outstanding delegator shares
	DelegatorDeductions sdk.Dec             // Delegator deductions from validator's delegators voting independently
	Vote                WeightedVoteOptions // Vote of the validator
}

// NewValidatorGovInfo creates a ValidatorGovInfo instance
func NewValidatorGovInfo(address sdk.ValAddress, bondedTokens sdk.Int, delegatorShares,
	delegatorDeductions sdk.Dec, options WeightedVoteOptions) ValidatorGovInfo {

	return ValidatorGovInfo{
		Address:             address,
		BondedTokens:        bondedTokens,
		DelegatorShares:     delegatorShares,
		DelegatorDeductions: delegatorDeductions,
		Vote:                options,
	}
}

// NewTallyResult creates a new TallyResult instance
func NewTallyResult(yes, abstain, no, noWithVeto sdk.Int) TallyResult {
	return TallyResult{
		YesCount:        yes.String(),
		AbstainCount:    abstain.String(),
		NoCount:         no.String(),
		NoWithVetoCount: noWithVeto.String(),
	}
}

// NewTallyResultFromMap creates a new TallyResult instance from a Option -> Dec map
func NewTallyResultFromMap(results map[VoteOption]sdk.Dec) TallyResult {
	return NewTallyResult(
		results[OptionYes].TruncateInt(),
		results[OptionAbstain].TruncateInt(),
		results[OptionNo].TruncateInt(),
		results[OptionNoWithVeto].TruncateInt(),
	)
}

// EmptyTallyResult returns an empty TallyResult.
func EmptyTallyResult() TallyResult {
	return NewTallyResult(sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt())
}

// Equals returns if two tally results are equal.
func (tr TallyResult) Equals(comp TallyResult) bool {
	return tr.YesCount == comp.YesCount &&
		tr.AbstainCount == comp.AbstainCount &&
		tr.NoCount == comp.NoCount &&
		tr.NoWithVetoCount == comp.NoWithVetoCount
}
