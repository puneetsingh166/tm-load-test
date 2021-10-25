package crisis

import (
	"time"

	"github.com/onomyprotocol/tm-load-test/customclient/telemetry"
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
	"github.com/onomyprotocol/tm-load-test/customclient/x/crisis/keeper"
	"github.com/onomyprotocol/tm-load-test/customclient/x/crisis/types"
)

// check all registered invariants
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	if k.InvCheckPeriod() == 0 || ctx.BlockHeight()%int64(k.InvCheckPeriod()) != 0 {
		// skip running the invariant check
		return
	}
	k.AssertInvariants(ctx)
}
