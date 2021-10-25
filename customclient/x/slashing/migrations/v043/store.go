package v043

import (
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
	v043distribution "github.com/onomyprotocol/tm-load-test/customclient/x/distribution/migrations/v043"
	v040slashing "github.com/onomyprotocol/tm-load-test/customclient/x/slashing/migrations/v040"
)

// MigrateStore performs in-place store migrations from v0.40 to v0.43. The
// migration includes:
//
// - Change addresses to be length-prefixed.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey) error {
	store := ctx.KVStore(storeKey)
	v043distribution.MigratePrefixAddress(store, v040slashing.ValidatorSigningInfoKeyPrefix)
	v043distribution.MigratePrefixAddressBytes(store, v040slashing.ValidatorMissedBlockBitArrayKeyPrefix)
	v043distribution.MigratePrefixAddress(store, v040slashing.AddrPubkeyRelationKeyPrefix)

	return nil
}
