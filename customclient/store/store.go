package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/onomyprotocol/tm-load-test/customclient/store/cache"
	"github.com/onomyprotocol/tm-load-test/customclient/store/rootmulti"
	"github.com/onomyprotocol/tm-load-test/customclient/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
