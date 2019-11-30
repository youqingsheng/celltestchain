package cli

import (
	"fmt"

	"github.com/cell-network/cellchain/x/cell/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the nameservice module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	nameserviceQueryCmd.AddCommand(client.GetCommands(
		GetCmdQueryShareHolders(storeKey, cdc),
	)...)
	return nameserviceQueryCmd
}

// GetCmdNames queries a list of all names
func GetCmdQueryShareHolders(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "shares",
		Short: "shares",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			valAddr := args[0]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/shares/%s", queryRoute, valAddr), nil)
			if err != nil {
				fmt.Printf("could not get query shareholders\n %s", err.Error())
				return nil
			}

			var out types.Share
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
