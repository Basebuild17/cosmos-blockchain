# Registry CLI

This directory adds two small CLI helpers for the x/registry module:

- `tx.go` — transaction commands (`registry set <key> <value>`)
- `query.go` — query commands (`registry get <key>`, `registry list`)

Quick usage

- Build your CLI binaries as documented in GETTING_STARTED.md (example names used below):

  - tx (set a key/value):

    ```bash
    ./myblockchaincli tx registry set mykey myvalue --from <keyname> --chain-id <chain-id> --fees "100stake" -y
    ```

  - query (get a key):

    ```bash
    ./myblockchaincli query registry get mykey --node tcp://127.0.0.1:26657
    ```

  - query (list all):

    ```bash
    ./myblockchaincli query registry list --node tcp://127.0.0.1:26657
    ```

Notes about flags

- `--from <keyname>`: signer name from local keyring (required for tx)
- `--chain-id <id>`: your chain id (can also be in config)
- `--fees "100stake"`: transaction fees — adjust to your token denom
- `-y`: auto-confirm the transaction

Wiring the commands into the root CLI

To make the commands available in your application's root CLI binary, import the CLI package and register the module commands where your root command is assembled. Example (pseudo-code for your `cmd/<app>/root.go`):

```go
import (
    registrycli "github.com/Basebuild17/cosmos-blockchain/x/registry/client/cli"
)

func AddModuleCommands(rootCmd *cobra.Command) {
    // register tx commands
    rootCmd.AddCommand(registrycli.NewTxCmd())

    // register query commands
    rootCmd.AddCommand(registrycli.NewQueryCmd())
}
```

If you point me to the real file that builds the root command (typically under `cmd/` or `./cmd/<app>/root.go`), I can open another change to wire the commands into that file for you.
