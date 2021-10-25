package legacytx

import (
	"github.com/onomyprotocol/tm-load-test/customclient/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(StdTx{}, "cosmos-sdk/StdTx", nil)
}
