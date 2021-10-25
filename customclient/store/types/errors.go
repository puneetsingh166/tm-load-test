package types

import (
	sdkerrors "github.com/onomyprotocol/tm-load-test/customclient/types/errors"
)

const StoreCodespace = "store"

var (
	ErrInvalidProof = sdkerrors.Register(StoreCodespace, 2, "invalid proof")
)
