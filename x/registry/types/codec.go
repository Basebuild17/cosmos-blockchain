package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codeTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// ModuleCdc references the global x/registry module codec. Note, the codec should
// ONLY be used to marshal/unmarshal cache values, the codec used for state
// migration should always be used when working with core types.
var ModuleCdc = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
}

// RegisterLegacyAminoCodec registers the necessary x/registry interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetData{}, "registry/SetData", nil)
	cdc.RegisterConcrete(&MsgGetData{}, "registry/GetData", nil)
}

// RegisterInterfaces registers the x/registry interfaces types with the interface registry.
func RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSetData{},
		&MsgGetData{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
