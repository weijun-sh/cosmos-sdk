package client

import (
	govclient "github.com/weijun-sh/cosmos-sdk/x/gov/client"
	"github.com/weijun-sh/cosmos-sdk/x/upgrade/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal)
