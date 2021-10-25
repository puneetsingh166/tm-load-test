package testutil

import (
	"github.com/onomyprotocol/tm-load-test/customclient/testutil"
	clitestutil "github.com/onomyprotocol/tm-load-test/customclient/testutil/cli"
	"github.com/onomyprotocol/tm-load-test/customclient/testutil/network"
	"github.com/onomyprotocol/tm-load-test/customclient/x/authz/client/cli"
)

func ExecGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
