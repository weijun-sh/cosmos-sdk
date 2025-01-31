package keeper

import (
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	v046 "github.com/weijun-sh/cosmos-sdk/x/feegrant/migrations/v046"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v046.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
