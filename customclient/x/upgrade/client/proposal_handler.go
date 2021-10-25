package client

import (
	govclient "github.com/onomyprotocol/tm-load-test/customclient/x/gov/client"
	"github.com/onomyprotocol/tm-load-test/customclient/x/upgrade/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal)
