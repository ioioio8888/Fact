package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ioioio8888/factio/x/factio/internal/types"
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
		GetCmdGetFactList(storeKey, cdc),
		GetCmdGetAccountCoins(storeKey, cdc),
		GetCmdGetVotePower(storeKey, cdc),
		GetCmdGetVoteOnFactList(storeKey, cdc),
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

// GetCmdGetAccountCoins queries a list of coins that address has
func GetCmdGetAccountCoins(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-acc-coins [address]",
		Short: "get-acc-coins [address]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			address := args[0]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getAccCoins/%s", queryRoute, address), nil)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return nil
			}
			var out sdk.Coins
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdGetVotePower queries a votepower
func GetCmdGetVotePower(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-vote-power [address]",
		Short: "get-vote-power [address]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			address := args[0]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getVotePower/%s", queryRoute, address), nil)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return nil
			}
			var out types.VotePower
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdGetVoteOnFact queries a list of the fact that the address has voted
func GetCmdGetVoteOnFactList(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-voted-list [address]",
		Short: "get-voted-list [address]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			address := args[0]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getVoteList/%s", queryRoute, address), nil)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return nil
			}
			var out types.QueryResVotedList
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
