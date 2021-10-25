package types

import (
	sdkerrors "github.com/onomyprotocol/tm-load-test/customclient/types/errors"
)

// x/crisis module sentinel errors
var (
	ErrNoSender         = sdkerrors.Register(ModuleName, 2, "sender address is empty")
	ErrUnknownInvariant = sdkerrors.Register(ModuleName, 3, "unknown invariant")
)
