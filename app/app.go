package app

import (
	"io"

	wasmapp "github.com/CosmWasm/wasmd/app"
	wasmparams "github.com/CosmWasm/wasmd/app/params"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cometbft/cometbft/libs/log"
	dbm "github.com/cometbft/cometbft-db"
)

// MyBlockchainApp embeds the standard wasmd application so this chain gets
// CosmWasm's keeper, gas metering, permissions, messages, queries, and CLI.
type MyBlockchainApp struct {
	*wasmapp.WasmApp
}

// NewMyBlockchainApp creates a CosmWasm-enabled application.
func NewMyBlockchainApp(logger log.Logger, db dbm.DB, traceStore string, loadLatest bool) *MyBlockchainApp {
	app := wasmapp.NewWasmApp(
		logger,
		db,
		io.Discard,
		loadLatest,
		server.DefaultAppOptions{},
		[]wasmkeeper.Option{},
	)

	return &MyBlockchainApp{WasmApp: app}
}

// EncodingConfig is the encoding configuration used by wasmd.
type EncodingConfig = wasmparams.EncodingConfig

// MakeEncodingConfig returns the standard CosmWasm encoding configuration.
func MakeEncodingConfig() EncodingConfig {
	return wasmapp.MakeEncodingConfig()
}

// GetTxConfig returns the transaction configuration for signing and decoding.
func (app *MyBlockchainApp) GetTxConfig() client.TxConfig {
	return app.TxConfig()
}

// GetCodec returns the application codec.
func (app *MyBlockchainApp) GetCodec() codec.Codec {
	return app.AppCodec()
}
# VeilPay — private-payments prototype

VeilPay is a small, local **permissioned-ledger prototype** for exploring private
payments. A validator accepts transfers only from registered wallets, verifies a
wallet signature, prevents replay, and commits accepted transfers into a
hash-linked chain.

It is intentionally a learning/MVP implementation, not a cryptocurrency to
deploy with real funds.

## Run the demo

Requires Python 3.11+ and no external packages.

```bash
python3 -m unittest discover -s tests -v
python3 demo.py
```

The demo creates two wallets, funds Alice in the genesis block, sends Bob a
payment with an encrypted memo, then verifies the whole chain.

## Privacy boundary

Wallet secrets never leave the client and the optional payment memo is encrypted
for the recipient. Validator membership is allow-listed.

Amounts, sender addresses, and recipient addresses are still visible to
validators. **This is not confidential-transactions or zero-knowledge privacy.**
Those guarantees need a reviewed protocol and audited cryptographic library;
they must not be improvised in application code.

## Layout

- `veilpay.wallet` creates and signs with locally held secrets.
- `veilpay.ledger` validates transactions, prevents replay, and maintains blocks.
- `demo.py` is the smallest end-to-end example.

## Before real deployment

Replace the educational shared-secret wallet model with a vetted Ed25519 library,
store keys in an HSM or operating-system keychain, use mutually authenticated TLS
between nodes, persist state atomically, add Byzantine-fault tolerant consensus,
and complete an independent security review. For true payment privacy, adopt a
reviewed confidential-transactions or zero-knowledge protocol.
