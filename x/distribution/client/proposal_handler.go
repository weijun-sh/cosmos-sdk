package client

import (
	"github.com/weijun-sh/cosmos-sdk/x/distribution/client/cli"
	govclient "github.com/weijun-sh/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
