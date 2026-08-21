package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// x/registry module sentinel errors
var (
	ErrInvalidKey   = sdkerrors.Register(ModuleName, 1, "invalid key")
	ErrInvalidValue = sdkerrors.Register(ModuleName, 2, "invalid value")
	ErrKeyNotFound  = sdkerrors.Register(ModuleName, 3, "key not found")
)
