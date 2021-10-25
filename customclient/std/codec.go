package std

import (
	"github.com/onomyprotocol/tm-load-test/customclient/codec"
	"github.com/onomyprotocol/tm-load-test/customclient/codec/types"
	cryptocodec "github.com/onomyprotocol/tm-load-test/customclient/crypto/codec"
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
	txtypes "github.com/onomyprotocol/tm-load-test/customclient/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
