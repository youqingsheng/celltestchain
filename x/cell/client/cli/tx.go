package cli

import (
	"github.com/cell-network/cellchain/x/cell/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "cell transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	nameserviceTxCmd.AddCommand(client.PostCommands(
		GetCmdSetShareHolder(cdc),
		GetCmdDeleteShareHolder(cdc),
	)...)

	return nameserviceTxCmd
}

// GetCmdSetName is the CLI command for sending a SetName transaction
func GetCmdSetShareHolder(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set [delegator-addr] [rate]",
		Short: "set/add a delegator to share validator's commission",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }
			delegatorAddr, errs := sdk.AccAddressFromBech32(args[0])
			if errs != nil {
				return errs
			}
			rate, err2 := sdk.NewDecFromStr(args[1])
			if err2 != nil {
				return err2
			}
			msg := types.NewMsgSetShareHolder(sdk.ValAddress(cliCtx.GetFromAddress().Bytes()), delegatorAddr, rate)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteName is the CLI command for sending a DeleteName transaction
func GetCmdDeleteShareHolder(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [delegator-addr]",
		Short: "delete delegator that share your commissions",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			delegatorAddr, errs := sdk.AccAddressFromBech32(args[0])
			if errs != nil {
				return errs
			}

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteShareHolder(sdk.ValAddress(cliCtx.GetFromAddress().Bytes()), delegatorAddr)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
