package v046

import (
	"testing"

	"github.com/weijun-sh/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	"github.com/weijun-sh/cosmos-sdk/types/address"
	bank "github.com/weijun-sh/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
)

var granter = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
var grantee = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
var msgType = bank.SendAuthorization{}.MsgTypeURL()

func TestGrantkey(t *testing.T) {
	require := require.New(t)
	key := GrantStoreKey(grantee, granter, msgType)
	require.Len(key, len(GrantPrefix)+len(address.MustLengthPrefix(grantee))+len(address.MustLengthPrefix(granter))+len([]byte(msgType)))

	granter1, grantee1, msgType1 := ParseGrantKey(key[1:])
	require.Equal(granter, granter1)
	require.Equal(grantee, grantee1)
	require.Equal(msgType1, msgType)
}
