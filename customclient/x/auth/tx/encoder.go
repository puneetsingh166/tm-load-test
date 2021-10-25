package tx

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/onomyprotocol/tm-load-test/customclient/codec"
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
	txtypes "github.com/onomyprotocol/tm-load-test/customclient/types/tx"
)

// DefaultTxEncoder returns a default protobuf TxEncoder using the provided Marshaler
func DefaultTxEncoder() sdk.TxEncoder {
	return func(tx sdk.Tx) ([]byte, error) {
		txWrapper, ok := tx.(*wrapper)
		if !ok {
			return nil, fmt.Errorf("expected %T, got %T", &wrapper{}, tx)
		}

		raw := &txtypes.TxRaw{
			BodyBytes:     txWrapper.getBodyBytes(),
			AuthInfoBytes: txWrapper.getAuthInfoBytes(),
			Signatures:    txWrapper.tx.Signatures,
		}

		return proto.Marshal(raw)
	}
}

// DefaultJSONTxEncoder returns a default protobuf JSON TxEncoder using the provided Marshaler.
func DefaultJSONTxEncoder(cdc codec.ProtoCodecMarshaler) sdk.TxEncoder {
	return func(tx sdk.Tx) ([]byte, error) {
		txWrapper, ok := tx.(*wrapper)
		if ok {
			return cdc.MarshalJSON(txWrapper.tx)
		}

		protoTx, ok := tx.(*txtypes.Tx)
		if ok {
			return cdc.MarshalJSON(protoTx)
		}

		return nil, fmt.Errorf("expected %T, got %T", &wrapper{}, tx)

	}
}
