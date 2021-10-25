package keeper_test

import (
	"github.com/onomyprotocol/tm-load-test/customclient/codec"
	"github.com/onomyprotocol/tm-load-test/customclient/simapp"
	"github.com/onomyprotocol/tm-load-test/customclient/testutil"
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
	paramskeeper "github.com/onomyprotocol/tm-load-test/customclient/x/params/keeper"
)

func testComponents() (*codec.LegacyAmino, sdk.Context, sdk.StoreKey, sdk.StoreKey, paramskeeper.Keeper) {
	marshaler := simapp.MakeTestEncodingConfig().Codec
	legacyAmino := createTestCodec()
	mkey := sdk.NewKVStoreKey("test")
	tkey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(mkey, tkey)
	keeper := paramskeeper.NewKeeper(marshaler, legacyAmino, mkey, tkey)

	return legacyAmino, ctx, mkey, tkey, keeper
}

type invalid struct{}

type s struct {
	I int
}

func createTestCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cdc.RegisterConcrete(s{}, "test/s", nil)
	cdc.RegisterConcrete(invalid{}, "test/invalid", nil)
	return cdc
}
