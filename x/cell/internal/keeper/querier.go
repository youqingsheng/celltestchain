package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the nameservice Querier
const (
	QueryShares = "shares"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryShares:
			return queryShares(ctx, path, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown cell query endpoint")
		}
	}
}

// nolint: unparam
func queryShares(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {

	fmt.Printf("In Query Shares Keeper [%s]", path[1])

	valAddr, err := sdk.ValAddressFromBech32(path[1])
	if err != nil {
		print("Invalidate Address")
	}
	share := keeper.GetShare(ctx, valAddr)

	res, err := codec.MarshalJSONIndent(keeper.cdc, share)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
