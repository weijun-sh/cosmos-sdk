package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/weijun-sh/cosmos-sdk/codec"
	cryptoAmino "github.com/weijun-sh/cosmos-sdk/crypto/codec"
	"github.com/weijun-sh/cosmos-sdk/testutil/testdata"
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	"github.com/weijun-sh/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/weijun-sh/cosmos-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
