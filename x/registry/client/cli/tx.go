package cli

import (
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/Basebuild17/cosmos-blockchain/x/registry/types"
)

// NewTxCmd returns the CLI commands for registry transactions
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "registry",
		Short: "Registry module transactions",
	}
	txCmd.AddCommand(CmdSetData())
	return txCmd
}

// CmdSetData creates a set-data tx
func CmdSetData() *cobra.Command {
	return &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set a key/value in the registry",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetData(clientCtx.GetFromAddress().String(), args[0], args[1])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
