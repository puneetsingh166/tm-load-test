package v039

import sdk "github.com/onomyprotocol/tm-load-test/customclient/types"

const (
	ModuleName = "crisis"
)

type (
	GenesisState struct {
		ConstantFee sdk.Coin `json:"constant_fee" yaml:"constant_fee"`
	}
)
