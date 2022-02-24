package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/weijun-sh/cosmos-sdk/codec"
	codectypes "github.com/weijun-sh/cosmos-sdk/codec/types"
	"github.com/weijun-sh/cosmos-sdk/std"
	"github.com/weijun-sh/cosmos-sdk/testutil/testdata"
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	"github.com/weijun-sh/cosmos-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
