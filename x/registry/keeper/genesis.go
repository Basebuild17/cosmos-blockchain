package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Basebuild17/cosmos-blockchain/x/registry/types"
)

// InitGenesis initializes the module genesis state
func (k Keeper) InitGenesis(ctx sdk.Context, genesisState types.GenesisState) {
	// Initialize all data from genesis
	for _, data := range genesisState.DataList {
		k.SetData(ctx, data.Key, data.Value)
	}
}

// ExportGenesis returns the module's genesis state
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.DataList = k.GetAllData(ctx)
	return *genesis
}
