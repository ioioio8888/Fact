package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/sdk-tutorials/factio/x/factio/internal/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	factioQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the factio module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	factioQueryCmd.AddCommand(client.GetCommands(
		GetCmdGetFact(storeKey, cdc),
		GetCmdGetAddressDelegation(storeKey, cdc),
		GetCmdGetFactList(storeKey, cdc),
	)...)
	return factioQueryCmd
}

// GetCmdGetFact queries information about a name
func GetCmdGetFact(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-fact [title]",
		Short: "get-fact title",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			title := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getFact/%s", queryRoute, title), nil)
			if err != nil {
				fmt.Printf("could not get the title - %s \n", title)
				return nil
			}

			var out types.Fact
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdGetAddressdelegation queries a list of fact that the address has been delegated to
func GetCmdGetAddressDelegation(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-addrdele [address]",
		Short: "get-addrdele address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			address := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getAddressDelegation/%s", queryRoute, address), nil)
			if err != nil {
				fmt.Printf("could not get the title - %s \n", address)
				return nil
			}

			var out types.QueryResFactDelegationList
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdGetFactList queries a list of fact
func GetCmdGetFactList(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-factlist",
		Short: "get-factlist",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getFactList", queryRoute), nil)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return nil
			}

			var out types.QueryResFactList
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
