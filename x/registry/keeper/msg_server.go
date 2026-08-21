package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Basebuild17/cosmos-blockchain/x/registry/types"
)

type msgServer struct {
	Keeper Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// SetData handles the SetData message
func (k msgServer) SetData(goCtx context.Context, msg *types.MsgSetData) (*types.MsgSetDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Store the data
	k.Keeper.SetData(ctx, msg.Key, msg.Value)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSetData,
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyKey, msg.Key),
			sdk.NewAttribute(types.AttributeKeyValue, msg.Value),
		),
	)

	return &types.MsgSetDataResponse{}, nil
}

// GetData handles the GetData message
func (k msgServer) GetData(goCtx context.Context, msg *types.MsgGetData) (*types.MsgGetDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Retrieve the data
	value, found := k.Keeper.GetData(ctx, msg.Key)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrKeyNotFound, msg.Key)
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeGetData,
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyKey, msg.Key),
			sdk.NewAttribute(types.AttributeKeyValue, value),
		),
	)

	return &types.MsgGetDataResponse{Value: value}, nil
}
