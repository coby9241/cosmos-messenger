package cosmosmessenger_test

import (
	"testing"

	keepertest "cosmos-messenger/testutil/keeper"
	"cosmos-messenger/testutil/nullify"
	"cosmos-messenger/x/cosmosmessenger"
	"cosmos-messenger/x/cosmosmessenger/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosmessengerKeeper(t)
	cosmosmessenger.InitGenesis(ctx, *k, genesisState)
	got := cosmosmessenger.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
