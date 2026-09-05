package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/Basebuild17/cosmos-blockchain/app"
	registrycli "github.com/Basebuild17/cosmos-blockchain/x/registry/client/cli"
)

// NewRootCmd creates the root CLI command for the myblockchain daemon binary.
func NewRootCmd() *cobra.Command {
	// Build encoding config from app package
	enc := app.MakeEncodingConfig()

	clientCtx := client.Context{}.
		WithCodec(enc.Codec).
		WithLegacyAmino(enc.Amino).
		WithTxConfig(enc.TxConfig)

	rootCmd := &cobra.Command{
		Use:   "myblockaind",
		Short: "Daemon for MyBlockchain",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return client.SetCmdClientContextHandler(clientCtx, cmd)
		},
	}

	// tx and query root groups
	txRoot := &cobra.Command{Use: "tx", Short: "Transactions commands"}
	queryRoot := &cobra.Command{Use: "query", Short: "Querying commands"}

	// Register module-level commands under the groups
	txRoot.AddCommand(registrycli.NewTxCmd())
	queryRoot.AddCommand(registrycli.NewQueryCmd())

	// Add SDK standard flags to groups
	flags.AddTxFlagsToCmd(txRoot)
	flags.AddQueryFlagsToCmd(queryRoot)

	rootCmd.AddCommand(txRoot, queryRoot)

	return rootCmd
}
