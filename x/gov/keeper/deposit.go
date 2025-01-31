package keeper

import (
	"fmt"

	sdk "github.com/weijun-sh/cosmos-sdk/types"
	sdkerrors "github.com/weijun-sh/cosmos-sdk/types/errors"
	"github.com/weijun-sh/cosmos-sdk/x/gov/types"
	"github.com/weijun-sh/cosmos-sdk/x/gov/types/v1beta2"
)

// GetDeposit gets the deposit of a specific depositor on a specific proposal
func (keeper Keeper) GetDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) (deposit v1beta2.Deposit, found bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.DepositKey(proposalID, depositorAddr))
	if bz == nil {
		return deposit, false
	}

	keeper.cdc.MustUnmarshal(bz, &deposit)

	return deposit, true
}

// SetDeposit sets a Deposit to the gov store
func (keeper Keeper) SetDeposit(ctx sdk.Context, deposit v1beta2.Deposit) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshal(&deposit)
	depositor, err := sdk.AccAddressFromBech32(deposit.Depositor)
	if err != nil {
		panic(err)
	}

	store.Set(types.DepositKey(deposit.ProposalId, depositor), bz)
}

// GetAllDeposits returns all the deposits from the store
func (keeper Keeper) GetAllDeposits(ctx sdk.Context) (deposits v1beta2.Deposits) {
	keeper.IterateAllDeposits(ctx, func(deposit v1beta2.Deposit) bool {
		deposits = append(deposits, &deposit)
		return false
	})

	return
}

// GetDeposits returns all the deposits from a proposal
func (keeper Keeper) GetDeposits(ctx sdk.Context, proposalID uint64) (deposits v1beta2.Deposits) {
	keeper.IterateDeposits(ctx, proposalID, func(deposit v1beta2.Deposit) bool {
		deposits = append(deposits, &deposit)
		return false
	})

	return
}

// DeleteAndBurnDeposits deletes and burn all the deposits on a specific proposal.
func (keeper Keeper) DeleteAndBurnDeposits(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)

	keeper.IterateDeposits(ctx, proposalID, func(deposit v1beta2.Deposit) bool {
		err := keeper.bankKeeper.BurnCoins(ctx, types.ModuleName, deposit.Amount)
		if err != nil {
			panic(err)
		}

		depositor, err := sdk.AccAddressFromBech32(deposit.Depositor)
		if err != nil {
			panic(err)
		}
		store.Delete(types.DepositKey(proposalID, depositor))
		return false
	})
}

// IterateAllDeposits iterates over the all the stored deposits and performs a callback function
func (keeper Keeper) IterateAllDeposits(ctx sdk.Context, cb func(deposit v1beta2.Deposit) (stop bool)) {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DepositsKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var deposit v1beta2.Deposit

		keeper.cdc.MustUnmarshal(iterator.Value(), &deposit)

		if cb(deposit) {
			break
		}
	}
}

// IterateDeposits iterates over the all the proposals deposits and performs a callback function
func (keeper Keeper) IterateDeposits(ctx sdk.Context, proposalID uint64, cb func(deposit v1beta2.Deposit) (stop bool)) {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DepositsKey(proposalID))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var deposit v1beta2.Deposit

		keeper.cdc.MustUnmarshal(iterator.Value(), &deposit)

		if cb(deposit) {
			break
		}
	}
}

// AddDeposit adds or updates a deposit of a specific depositor on a specific proposal
// Activates voting period when appropriate
func (keeper Keeper) AddDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress, depositAmount sdk.Coins) (bool, error) {
	// Checks to see if proposal exists
	proposal, ok := keeper.GetProposal(ctx, proposalID)
	if !ok {
		return false, sdkerrors.Wrapf(types.ErrUnknownProposal, "%d", proposalID)
	}

	// Check if proposal is still depositable
	if (proposal.Status != v1beta2.StatusDepositPeriod) && (proposal.Status != v1beta2.StatusVotingPeriod) {
		return false, sdkerrors.Wrapf(types.ErrInactiveProposal, "%d", proposalID)
	}

	// update the governance module's account coins pool
	err := keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, depositorAddr, types.ModuleName, depositAmount)
	if err != nil {
		return false, err
	}

	// Update proposal
	proposal.TotalDeposit = sdk.NewCoins(proposal.TotalDeposit...).Add(depositAmount...)
	keeper.SetProposal(ctx, proposal)

	// Check if deposit has provided sufficient total funds to transition the proposal into the voting period
	activatedVotingPeriod := false

	if proposal.Status == v1beta2.StatusDepositPeriod && sdk.NewCoins(proposal.TotalDeposit...).IsAllGTE(keeper.GetDepositParams(ctx).MinDeposit) {
		keeper.ActivateVotingPeriod(ctx, proposal)

		activatedVotingPeriod = true
	}

	// Add or update deposit object
	deposit, found := keeper.GetDeposit(ctx, proposalID, depositorAddr)

	if found {
		deposit.Amount = sdk.NewCoins(deposit.Amount...).Add(depositAmount...)
	} else {
		deposit = v1beta2.NewDeposit(proposalID, depositorAddr, depositAmount)
	}

	// called when deposit has been added to a proposal, however the proposal may not be active
	keeper.AfterProposalDeposit(ctx, proposalID, depositorAddr)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeProposalDeposit,
			sdk.NewAttribute(sdk.AttributeKeyAmount, depositAmount.String()),
			sdk.NewAttribute(types.AttributeKeyProposalID, fmt.Sprintf("%d", proposalID)),
		),
	)

	keeper.SetDeposit(ctx, deposit)

	return activatedVotingPeriod, nil
}

// RefundAndDeleteDeposits refunds and deletes all the deposits on a specific proposal.
func (keeper Keeper) RefundAndDeleteDeposits(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)

	keeper.IterateDeposits(ctx, proposalID, func(deposit v1beta2.Deposit) bool {
		depositor, err := sdk.AccAddressFromBech32(deposit.Depositor)
		if err != nil {
			panic(err)
		}

		err = keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositor, deposit.Amount)
		if err != nil {
			panic(err)
		}

		store.Delete(types.DepositKey(proposalID, depositor))
		return false
	})
}
