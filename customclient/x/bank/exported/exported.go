package exported

import (
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
)

// GenesisBalance defines a genesis balance interface that allows for account
// address and balance retrieval.
type GenesisBalance interface {
	GetAddress() sdk.AccAddress
	GetCoins() sdk.Coins
}
