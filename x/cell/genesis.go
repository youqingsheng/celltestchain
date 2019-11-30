package cell

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Shares []Share `json:"shares"`
}

func NewGenesisState(whoIsRecords []Share) GenesisState {
	return GenesisState{Shares: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Shares {
		if record.Owner == nil {
			return fmt.Errorf("invalid Shares: Value: %s. Error: Missing Owner", record.Owner)
		}
		if record.ShareHolders == nil {
			return fmt.Errorf("invalid Shares: Owner: %s. Error: Missing Shareholders", record.ShareHolders)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Shares: []Share{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.Shares {
		keeper.SetShare(ctx, record.Owner, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Share
	iterator := k.GetValidatorsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := sdk.ValAddress(iterator.Key())
		share := k.GetShare(ctx, name)
		records = append(records, share)

	}
	return GenesisState{Shares: records}
}
