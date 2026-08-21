package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Basebuild17/cosmos-blockchain/x/registry/types"
)

// Querier is used as a namespace for query methods on the registry keeper.
type Querier struct {
	Keeper Keeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &Querier{Keeper: keeper}
}

var _ types.QueryServer = Querier{}

// GetData returns the data for a given key
func (q Querier) GetData(goCtx context.Context, req *types.QueryGetDataRequest) (*types.QueryGetDataResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	value, found := q.Keeper.GetData(ctx, req.Key)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrKeyNotFound, req.Key)
	}

	return &types.QueryGetDataResponse{Value: value}, nil
}

// ListData returns all data stored in the registry
func (q Querier) ListData(goCtx context.Context, req *types.QueryListDataRequest) (*types.QueryListDataResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	dataList := q.Keeper.GetAllData(ctx)

	return &types.QueryListDataResponse{Data: dataList}, nil
}
