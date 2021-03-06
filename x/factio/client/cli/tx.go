package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/ioioio8888/factio/x/factio/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	factioTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "factio transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	factioTxCmd.AddCommand(client.PostCommands(
		GetCmdCreateFact(cdc),
		GetCmdEditFact(cdc),
		// GetCmdDelegateFact(cdc),
		// GetCmdUnDelegateFact(cdc),
		GetCmdVoteFact(cdc),
		GetCmdUnVoteFact(cdc),
	)...)

	return factioTxCmd
}

// GetCmdCreateFact is the CLI command for sending a CreateFact transaction
func GetCmdCreateFact(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-fact [title] [time] [place] [description]",
		Short: "create a new fact if it is not exist",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			time, terr := strconv.ParseInt(args[1], 10, 64)
			if terr != nil {
				return terr
			}

			coins, err := sdk.ParseCoins("1factcoin")
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateFact(args[0], cliCtx.GetFromAddress(), coins, time, args[2], args[3])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdEditFact is the CLI command for sending a EditFact transaction
func GetCmdEditFact(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "edit-fact [title] [time] [place] [description]",
		Short: "set the content of a fact that you own",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			time, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}
			msg := types.NewMsgEditFact(args[0], cliCtx.GetFromAddress(), time, args[2], args[3])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdvoteFact is the CLI command for sending a UnVoteFact transaction
func GetCmdVoteFact(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vote-fact [title] [stance]",
		Short: "vote on a fact, stance: upvote/downvote ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			var stance bool
			if args[1] == "upvote" {
				stance = true
			} else if args[1] == "downvote" {
				stance = false
			} else {
				return types.ErrFactDoesNotExist("Invalid stance")
			}
			msg := types.NewMsgVoteFact(args[0], cliCtx.GetFromAddress(), stance)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdunvoteFact is the CLI command for sending a UnVoteFact transaction
func GetCmdUnVoteFact(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "unvote-fact [title]",
		Short: "unvote a fact ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgUnVoteFact(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
