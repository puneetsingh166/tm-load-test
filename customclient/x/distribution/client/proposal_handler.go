package client

import (
	"github.com/onomyprotocol/tm-load-test/customclient/x/distribution/client/cli"
	govclient "github.com/onomyprotocol/tm-load-test/customclient/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
