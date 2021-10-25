package distribution

import (
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
	sdkerrors "github.com/onomyprotocol/tm-load-test/customclient/types/errors"
	"github.com/onomyprotocol/tm-load-test/customclient/x/distribution/keeper"
	"github.com/onomyprotocol/tm-load-test/customclient/x/distribution/types"
	govtypes "github.com/onomyprotocol/tm-load-test/customclient/x/gov/types"
)

func NewCommunityPoolSpendProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CommunityPoolSpendProposal:
			return keeper.HandleCommunityPoolSpendProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized distr proposal content type: %T", c)
		}
	}
}
