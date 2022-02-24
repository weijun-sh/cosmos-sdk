package main

import (
	"os"

	"github.com/weijun-sh/cosmos-sdk/cosmovisor"
	"github.com/weijun-sh/cosmos-sdk/cosmovisor/cmd/cosmovisor/cmd"
	cverrors "github.com/weijun-sh/cosmos-sdk/cosmovisor/errors"
)

func main() {
	cosmovisor.SetupLogging()
	if err := cmd.RunCosmovisorCommand(os.Args[1:]); err != nil {
		cverrors.LogErrors(cosmovisor.Logger, "", err)
		os.Exit(1)
	}
}
