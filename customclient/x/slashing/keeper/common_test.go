package keeper_test

import sdk "github.com/onomyprotocol/tm-load-test/customclient/types"

var (
	// The default power validators are initialized to have within tests
	InitTokens = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
)
