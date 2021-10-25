package simulation

import (
	"bytes"
	"fmt"

	"github.com/onomyprotocol/tm-load-test/customclient/codec"
	"github.com/onomyprotocol/tm-load-test/customclient/types/kv"
	"github.com/onomyprotocol/tm-load-test/customclient/x/authz"
	"github.com/onomyprotocol/tm-load-test/customclient/x/authz/keeper"
)

// NewDecodeStore returns a decoder function closure that umarshals the KVPair's
// Value to the corresponding authz type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], keeper.GrantKey):
			var grantA, grantB authz.Grant
			cdc.MustUnmarshal(kvA.Value, &grantA)
			cdc.MustUnmarshal(kvB.Value, &grantB)
			return fmt.Sprintf("%v\n%v", grantA, grantB)
		default:
			panic(fmt.Sprintf("invalid authz key %X", kvA.Key))
		}
	}
}
