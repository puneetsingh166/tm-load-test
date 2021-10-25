package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onomyprotocol/tm-load-test/customclient/codec"
	codectypes "github.com/onomyprotocol/tm-load-test/customclient/codec/types"
	"github.com/onomyprotocol/tm-load-test/customclient/types/module"
	simtypes "github.com/onomyprotocol/tm-load-test/customclient/types/simulation"
	"github.com/onomyprotocol/tm-load-test/customclient/x/evidence/simulation"
	"github.com/onomyprotocol/tm-load-test/customclient/x/evidence/types"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: 1000,
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var evidenceGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &evidenceGenesis)

	require.Len(t, evidenceGenesis.Evidence, 0)
}
