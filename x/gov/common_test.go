package gov_test

import (
	"bytes"
	"log"
	"sort"
	"testing"

	"github.com/weijun-sh/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/weijun-sh/cosmos-sdk/crypto/types"
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	authtypes "github.com/weijun-sh/cosmos-sdk/x/auth/types"
	"github.com/weijun-sh/cosmos-sdk/x/gov/types"
	"github.com/weijun-sh/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/weijun-sh/cosmos-sdk/x/gov/types/v1beta2"
	stakingtypes "github.com/weijun-sh/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

var (
	valTokens           = sdk.TokensFromConsensusPower(42, sdk.DefaultPowerReduction)
	TestProposal        = v1beta1.NewTextProposal("Test", "description")
	TestDescription     = stakingtypes.NewDescription("T", "E", "S", "T", "Z")
	TestCommissionRates = stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())
)

// mkTestLegacyContent creates a MsgExecLegacyContent for testing purposes.
func mkTestLegacyContent(t *testing.T) *v1beta2.MsgExecLegacyContent {
	msgContent, err := v1beta2.NewLegacyContent(TestProposal, authtypes.NewModuleAddress(types.ModuleName).String())
	require.NoError(t, err)

	return msgContent
}

// SortAddresses - Sorts Addresses
func SortAddresses(addrs []sdk.AccAddress) {
	byteAddrs := make([][]byte, len(addrs))

	for i, addr := range addrs {
		byteAddrs[i] = addr.Bytes()
	}

	SortByteArrays(byteAddrs)

	for i, byteAddr := range byteAddrs {
		addrs[i] = byteAddr
	}
}

// implement `Interface` in sort package.
type sortByteArrays [][]byte

func (b sortByteArrays) Len() int {
	return len(b)
}

func (b sortByteArrays) Less(i, j int) bool {
	// bytes package already implements Comparable for []byte.
	switch bytes.Compare(b[i], b[j]) {
	case -1:
		return true
	case 0, 1:
		return false
	default:
		log.Panic("not fail-able with `bytes.Comparable` bounded [-1, 1].")
		return false
	}
}

func (b sortByteArrays) Swap(i, j int) {
	b[j], b[i] = b[i], b[j]
}

// SortByteArrays - sorts the provided byte array
func SortByteArrays(src [][]byte) [][]byte {
	sorted := sortByteArrays(src)
	sort.Sort(sorted)
	return sorted
}

const contextKeyBadProposal = "contextKeyBadProposal"

var (
	pubkeys = []cryptotypes.PubKey{
		ed25519.GenPrivKey().PubKey(),
		ed25519.GenPrivKey().PubKey(),
		ed25519.GenPrivKey().PubKey(),
	}
)
