package simulation

import (
	"math/rand"
	"time"

	codectypes "github.com/weijun-sh/cosmos-sdk/codec/types"
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	"github.com/weijun-sh/cosmos-sdk/types/module"
	simtypes "github.com/weijun-sh/cosmos-sdk/types/simulation"
	"github.com/weijun-sh/cosmos-sdk/x/authz"
	banktypes "github.com/weijun-sh/cosmos-sdk/x/bank/types"
	"github.com/weijun-sh/cosmos-sdk/x/gov/types/v1beta2"
)

// genGrant returns a slice of authorization grants.
func genGrant(r *rand.Rand, accounts []simtypes.Account, genT time.Time) []authz.GrantAuthorization {
	authorizations := make([]authz.GrantAuthorization, len(accounts)-1)
	for i := 0; i < len(accounts)-1; i++ {
		granter := accounts[i]
		grantee := accounts[i+1]
		authorizations[i] = authz.GrantAuthorization{
			Granter:       granter.Address.String(),
			Grantee:       grantee.Address.String(),
			Authorization: generateRandomGrant(r),
			Expiration:    genT.AddDate(1, 0, 0),
		}
	}

	return authorizations
}

func generateRandomGrant(r *rand.Rand) *codectypes.Any {
	authorizations := make([]*codectypes.Any, 2)
	authorizations[0] = newAnyAuthorization(banktypes.NewSendAuthorization(sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000)))))
	authorizations[1] = newAnyAuthorization(authz.NewGenericAuthorization(sdk.MsgTypeURL(&v1beta2.MsgSubmitProposal{})))

	return authorizations[r.Intn(len(authorizations))]
}

func newAnyAuthorization(a authz.Authorization) *codectypes.Any {
	any, err := codectypes.NewAnyWithValue(a)
	if err != nil {
		panic(err)
	}

	return any
}

// RandomizedGenState generates a random GenesisState for authz.
func RandomizedGenState(simState *module.SimulationState) {
	var grants []authz.GrantAuthorization
	simState.AppParams.GetOrGenerate(
		simState.Cdc, "authz", &grants, simState.Rand,
		func(r *rand.Rand) { grants = genGrant(r, simState.Accounts, simState.GenTimestamp) },
	)

	authzGrantsGenesis := authz.NewGenesisState(grants)

	simState.GenState[authz.ModuleName] = simState.Cdc.MustMarshalJSON(authzGrantsGenesis)
}
