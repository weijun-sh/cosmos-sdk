package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/weijun-sh/cosmos-sdk/simapp"
	"github.com/weijun-sh/cosmos-sdk/testutil/testdata"
	"github.com/weijun-sh/cosmos-sdk/x/auth/middleware"
)

func TestRegisterMsgService(t *testing.T) {
	// Create an encoding config that doesn't register testdata Msg services.
	encCfg := simapp.MakeTestEncodingConfig()
	msr := middleware.NewMsgServiceRouter(encCfg.InterfaceRegistry)
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			msr,
			testdata.MsgServerImpl{},
		)
	})

	// Register testdata Msg services, and rerun `RegisterService`.
	testdata.RegisterInterfaces(encCfg.InterfaceRegistry)
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			msr,
			testdata.MsgServerImpl{},
		)
	})
}

func TestRegisterMsgServiceTwice(t *testing.T) {
	// Setup baseapp.
	encCfg := simapp.MakeTestEncodingConfig()
	msr := middleware.NewMsgServiceRouter(encCfg.InterfaceRegistry)
	testdata.RegisterInterfaces(encCfg.InterfaceRegistry)

	// First time registering service shouldn't panic.
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			msr,
			testdata.MsgServerImpl{},
		)
	})

	// Second time should panic.
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			msr,
			testdata.MsgServerImpl{},
		)
	})
}
