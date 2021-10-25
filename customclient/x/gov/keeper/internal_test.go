package keeper

import "github.com/onomyprotocol/tm-load-test/customclient/x/gov/types"

// UnsafeSetHooks updates the gov keeper's hooks, overriding any potential
// pre-existing hooks.
// WARNING: this function should only be used in tests.
func UnsafeSetHooks(k *Keeper, h types.GovHooks) {
	k.hooks = h
}
