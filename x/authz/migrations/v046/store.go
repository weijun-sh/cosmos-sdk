package v046

import (
	"github.com/weijun-sh/cosmos-sdk/codec"
	"github.com/weijun-sh/cosmos-sdk/internal/conv"
	"github.com/weijun-sh/cosmos-sdk/store/prefix"
	storetypes "github.com/weijun-sh/cosmos-sdk/store/types"
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	"github.com/weijun-sh/cosmos-sdk/x/authz"
)

// MigrateStore performs in-place store migrations from v0.45 to v0.46. The
// migration includes:
//
// - pruning expired authorizations
// - create secondary index for pruning expired authorizations
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	err := addExpiredGrantsIndex(ctx, store, cdc)
	if err != nil {
		return err
	}

	return nil
}

func addExpiredGrantsIndex(ctx sdk.Context, store storetypes.KVStore, cdc codec.BinaryCodec) error {
	grantsStore := prefix.NewStore(store, GrantPrefix)

	grantsIter := grantsStore.Iterator(nil, nil)
	defer grantsIter.Close()

	queueItems := make(map[string][]string)
	for ; grantsIter.Valid(); grantsIter.Next() {
		var grant authz.Grant
		bz := grantsIter.Value()
		if err := cdc.Unmarshal(bz, &grant); err != nil {
			return err
		}

		// delete expired authorization
		if grant.Expiration.Before(ctx.BlockTime()) {
			grantsStore.Delete(grantsIter.Key())
		} else {
			granter, grantee, msgType := ParseGrantKey(grantsIter.Key())
			key := GrantQueueKey(grant.Expiration, granter, grantee)

			queueItem, ok := queueItems[conv.UnsafeBytesToStr(key)]
			if !ok {
				queueItems[string(key)] = []string{msgType}
			} else {
				queueItem = append(queueItem, msgType)
				queueItems[string(key)] = queueItem
			}
		}
	}

	for key, v := range queueItems {
		bz, err := cdc.Marshal(&authz.GrantQueueItem{
			MsgTypeUrls: v,
		})
		if err != nil {
			return err
		}
		store.Set(conv.UnsafeStrToBytes(key), bz)
	}

	return nil
}
