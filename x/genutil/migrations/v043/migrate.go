package v043

import (
	"github.com/weijun-sh/cosmos-sdk/client"
	v040bank "github.com/weijun-sh/cosmos-sdk/x/bank/migrations/v040"
	v043bank "github.com/weijun-sh/cosmos-sdk/x/bank/migrations/v043"
	bank "github.com/weijun-sh/cosmos-sdk/x/bank/types"
	"github.com/weijun-sh/cosmos-sdk/x/genutil/types"
	v040gov "github.com/weijun-sh/cosmos-sdk/x/gov/migrations/v040"
	v043gov "github.com/weijun-sh/cosmos-sdk/x/gov/migrations/v043"
	gov "github.com/weijun-sh/cosmos-sdk/x/gov/types/v1beta1"
)

// Migrate migrates exported state from v0.40 to a v0.43 genesis state.
func Migrate(appState types.AppMap, clientCtx client.Context) types.AppMap {
	// Migrate x/gov.
	if appState[v040gov.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var oldGovState gov.GenesisState
		clientCtx.Codec.MustUnmarshalJSON(appState[v040gov.ModuleName], &oldGovState)

		// delete deprecated x/gov genesis state
		delete(appState, v040gov.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v043gov.ModuleName] = clientCtx.Codec.MustMarshalJSON(v043gov.MigrateJSON(&oldGovState))
	}

	if appState[v040bank.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var oldBankState bank.GenesisState
		clientCtx.Codec.MustUnmarshalJSON(appState[v040bank.ModuleName], &oldBankState)

		// delete deprecated x/bank genesis state
		delete(appState, v040bank.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v043bank.ModuleName] = clientCtx.Codec.MustMarshalJSON(v043bank.MigrateJSON(&oldBankState))
	}

	return appState
}
