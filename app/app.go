package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codeTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"

	"github.com/Basebuild17/cosmos-blockchain/x/registry"
	registrytypes "github.com/Basebuild17/cosmos-blockchain/x/registry/types"
)

// MyBlockchainApp extends an ABCI application, but with most of its logic to
// live in the ~/ modules.
type MyBlockchainApp struct {
	*runtime.App
	cdc *codec.LegacyCodec

	registry registry.AppModule
}

// NewMyBlockchainApp returns a reference to an initialized blockchain app
func NewMyBlockchainApp(
	logger log.Logger,
	db tmdb.DB,
	traceStore string,
	loadLatest bool,
) *MyBlockchainApp {
	// Create the application codec
	encodingConfig := MakeEncodingConfig()

	app := &MyBlockchainApp{
		cdc: encodingConfig.Codec.(*codec.LegacyCodec),
	}

	// Initialize the runtime app
	app.App = runtime.NewApp(
		"myblockchain",
		logger,
		db,
		encodingConfig.TxConfig,
		runtime.DefaultAppOptions{
			TxDecoder: encodingConfig.TxConfig.TxDecoder(),
		},
	)

	// Register modules
	app.registry = registry.NewAppModule()

	// Return app
	return app
}

// EncodingConfig specifies the concrete encoding types to use for a given app.
type EncodingConfig struct {
	InterfaceRegistry codecTypes.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyCodec
}

// MakeEncodingConfig creates the SerializationConfig for this app
func MakeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codecTypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	txCfg := authtypes.StdTxConfig{Cdc: amino}

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

// GetTxConfig returns the TxConfig
func (app *MyBlockchainApp) GetTxConfig() client.TxConfig {
	return app.App.TxConfig
}

// GetCodec returns the codec
func (app *MyBlockchainApp) GetCodec() codec.Codec {
	return app.cdc
}
