package v040

import (
	v039auth "github.com/weijun-sh/cosmos-sdk/x/auth/migrations/v039"
	v036supply "github.com/weijun-sh/cosmos-sdk/x/bank/migrations/v036"
	v038bank "github.com/weijun-sh/cosmos-sdk/x/bank/migrations/v038"
	"github.com/weijun-sh/cosmos-sdk/x/bank/types"
)

// Migrate accepts exported v0.39 x/auth and v0.38 x/bank genesis state and
// migrates it to v0.40 x/bank genesis state. The migration includes:
//
// - Moving balances from x/auth to x/bank genesis state.
// - Moving supply from x/supply to x/bank genesis state.
// - Re-encode in v0.40 GenesisState.
func Migrate(
	bankGenState v038bank.GenesisState,
	authGenState v039auth.GenesisState,
	supplyGenState v036supply.GenesisState,
) *types.GenesisState {
	balances := make([]types.Balance, len(authGenState.Accounts))
	for i, acc := range authGenState.Accounts {
		balances[i] = types.Balance{
			Address: acc.GetAddress().String(),
			Coins:   acc.GetCoins(),
		}
	}

	return &types.GenesisState{
		Params: types.Params{
			SendEnabled:        []*types.SendEnabled{},
			DefaultSendEnabled: bankGenState.SendEnabled,
		},
		Balances:      balances,
		Supply:        supplyGenState.Supply,
		DenomMetadata: []types.Metadata{},
	}
}
