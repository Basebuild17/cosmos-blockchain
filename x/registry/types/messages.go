package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSetData = "set_data"
	TypeMsgGetData = "get_data"
)

// NewMsgSetData creates a new MsgSetData message
func NewMsgSetData(creator, key, value string) *MsgSetData {
	return &MsgSetData{
		Creator: creator,
		Key:     key,
		Value:   value,
	}
}

// Route implements the sdk.Msg interface
func (msg *MsgSetData) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface
func (msg *MsgSetData) Type() string {
	return TypeMsgSetData
}

// GetSigners implements the sdk.Msg interface
func (msg *MsgSetData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements the sdk.Msg interface
func (msg *MsgSetData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface
func (msg *MsgSetData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Key == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "key cannot be empty")
	}

	if msg.Value == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "value cannot be empty")
	}

	return nil
}

// NewMsgGetData creates a new MsgGetData message
func NewMsgGetData(creator, key string) *MsgGetData {
	return &MsgGetData{
		Creator: creator,
		Key:     key,
	}
}

// Route implements the sdk.Msg interface
func (msg *MsgGetData) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface
func (msg *MsgGetData) Type() string {
	return TypeMsgGetData
}

// GetSigners implements the sdk.Msg interface
func (msg *MsgGetData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements the sdk.Msg interface
func (msg *MsgGetData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface
func (msg *MsgGetData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Key == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "key cannot be empty")
	}

	return nil
}
