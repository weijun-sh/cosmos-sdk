package authz

import (
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	"github.com/weijun-sh/cosmos-sdk/x/authz/keeper"
)

// BeginBlocker is called at the begining of every block
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {

	// delete all the mature grants
	if err := keeper.DequeueAndDeleteExpiredGrants(ctx); err != nil {
		panic(err)
	}
}
