package lumenx_test

import (
	"testing"

	keepertest "github.com/cryptonetD/lumenx/testutil/keeper"
	"github.com/cryptonetD/lumenx/testutil/nullify"
	"github.com/cryptonetD/lumenx/x/lumenx"
	"github.com/cryptonetD/lumenx/x/lumenx/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LumenxKeeper(t)
	lumenx.InitGenesis(ctx, *k, genesisState)
	got := lumenx.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
