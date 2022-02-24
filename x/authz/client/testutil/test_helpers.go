package testutil

import (
	"github.com/weijun-sh/cosmos-sdk/testutil"
	clitestutil "github.com/weijun-sh/cosmos-sdk/testutil/cli"
	"github.com/weijun-sh/cosmos-sdk/testutil/network"
	"github.com/weijun-sh/cosmos-sdk/x/authz/client/cli"
)

func CreateGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
