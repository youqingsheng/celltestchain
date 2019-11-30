package cellkeeper

import (
	"github.com/cell-network/cellchain/x/cell"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
)

// Keeper of the distribution store
type Keeper struct {
	cellKeeper   cell.Keeper
	distrKeeper  distr.Keeper
	codespace    sdk.CodespaceType
	supplyKeeper types.SupplyKeeper
	mintKeeper   mint.Keeper
}

// NewKeeper creates a new distribution Keeper instance
func NewKeeper(ck cell.Keeper, mintKeeper mint.Keeper, dk distr.Keeper, supplyKeeper types.SupplyKeeper, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		cellKeeper:   ck,
		distrKeeper:  dk,
		supplyKeeper: supplyKeeper,
		codespace:    codespace,
		mintKeeper:   mintKeeper,
	}
}

func (k Keeper) GetDistrKeeper() distr.Keeper {
	return k.distrKeeper
}

// withdraw validator commission
func (k Keeper) WithdrawValidatorCommission(ctx sdk.Context, valAddr sdk.ValAddress) (sdk.Coins, sdk.Error) {
	// fetch validator accumulated commission
	accumCommission := k.distrKeeper.GetValidatorAccumulatedCommission(ctx, valAddr)
	if accumCommission.IsZero() {
		return nil, types.ErrNoValidatorCommission(k.codespace)
	}

	commission, remainder := accumCommission.TruncateDecimal()
	k.distrKeeper.SetValidatorAccumulatedCommission(ctx, valAddr, remainder) // leave remainder to withdraw later

	// update outstanding
	outstanding := k.distrKeeper.GetValidatorOutstandingRewards(ctx, valAddr)
	k.distrKeeper.SetValidatorOutstandingRewards(ctx, valAddr, outstanding.Sub(sdk.NewDecCoins(commission)))

	if !commission.IsZero() {

		params := k.mintKeeper.GetParams(ctx)

		accAddr := sdk.AccAddress(valAddr)
		withdrawAddr := k.distrKeeper.GetDelegatorWithdrawAddr(ctx, accAddr)
		share := k.cellKeeper.GetShare(ctx, valAddr)
		remainder := commission
		for _, h := range share.ShareHolders {

			shareCommision := sdk.NewCoins(sdk.NewCoin(params.MintDenom, h.Rate.MulInt(commission.AmountOf(params.MintDenom)).TruncateInt()))

			if !shareCommision.IsZero() {
				remainder = remainder.Sub(shareCommision)
				err := k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, h.Address, shareCommision)
				if err != nil {
					return nil, err
				}
			}
		}

		err := k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, withdrawAddr, remainder)
		if err != nil {
			return nil, err
		}

	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawCommission,
			sdk.NewAttribute(sdk.AttributeKeyAmount, commission.String()),
		),
	)

	return commission, nil
}
