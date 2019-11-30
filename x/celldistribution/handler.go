package celldistribution

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgSetWithdrawAddress:
			return handleMsgModifyWithdrawAddress(ctx, msg, k)

		case MsgWithdrawDelegatorReward:
			return handleMsgWithdrawDelegatorReward(ctx, msg, k)

		case MsgWithdrawValidatorCommission:
			return handleMsgWithdrawValidatorCommission(ctx, msg, k)

		default:
			errMsg := fmt.Sprintf("unrecognized distribution message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// These functions assume everything has been authenticated (ValidateBasic passed, and signatures checked)

func handleMsgModifyWithdrawAddress(ctx sdk.Context, msg MsgSetWithdrawAddress, k Keeper) sdk.Result {
	err := k.GetDistrKeeper().SetWithdrawAddr(ctx, msg.DelegatorAddress, msg.WithdrawAddress)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress.String()),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgWithdrawDelegatorReward(ctx sdk.Context, msg MsgWithdrawDelegatorReward, k Keeper) sdk.Result {
	_, err := k.GetDistrKeeper().WithdrawDelegationRewards(ctx, msg.DelegatorAddress, msg.ValidatorAddress)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress.String()),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgWithdrawValidatorCommission(ctx sdk.Context, msg MsgWithdrawValidatorCommission, k Keeper) sdk.Result {
	_, err := k.WithdrawValidatorCommission(ctx, msg.ValidatorAddress)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddress.String()),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}
