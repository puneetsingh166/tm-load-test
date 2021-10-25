package types

import (
	"github.com/onomyprotocol/tm-load-test/customclient/codec"
	cryptocodec "github.com/onomyprotocol/tm-load-test/customclient/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
