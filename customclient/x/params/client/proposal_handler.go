package client

import (
	govclient "github.com/onomyprotocol/tm-load-test/customclient/x/gov/client"
	"github.com/onomyprotocol/tm-load-test/customclient/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
