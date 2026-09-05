package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/Basebuild17/cosmos-blockchain/x/registry/types"
)

// NewQueryCmd returns the CLI commands for registry queries
func NewQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:   "registry",
		Short: "Registry module queries",
	}
	queryCmd.AddCommand(CmdGetData())
	queryCmd.AddCommand(CmdListData())
	return queryCmd
}

// CmdGetData queries a single key
func CmdGetData() *cobra.Command {
	return &cobra.Command{
		Use:   "get [key]",
		Short: "Get a value by key from the registry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetData(context.Background(), &types.QueryGetDataRequest{Key: args[0]})
			if err != nil {
				return err
			}

			fmt.Println(res.Value)
			return nil
		},
	}
}

// CmdListData returns all registry entries
func CmdListData() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all registry key/value pairs",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ListData(context.Background(), &types.QueryListDataRequest{})
			if err != nil {
				return err
			}

			for _, d := range res.Data {
				fmt.Printf("%s: %s\n", d.Key, d.Value)
			}
			return nil
		},
	}
}
