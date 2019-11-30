package cell

import (
	"fmt"

	"github.com/cell-network/cellchain/x/cell/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetShareHolder:
			return handleMsgSetShareHolder(ctx, keeper, msg)
		case MsgDeleteShareHolder:
			return handleMsgDeleteShareHolder(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized shareholder Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set name
func handleMsgSetShareHolder(ctx sdk.Context, keeper Keeper, msg MsgSetShareHolder) sdk.Result {

	err := keeper.AddShareHolder(ctx, msg.Validator, types.NewShareHolder(msg.Delegator, msg.Rate))
	if err != nil {
		return err.Result()
	}

	return sdk.Result{} // return
}

// Handle a message to delete name
func handleMsgDeleteShareHolder(ctx sdk.Context, keeper Keeper, msg MsgDeleteShareHolder) sdk.Result {

	if !keeper.IsValidatorPresent(ctx, msg.Validator) {
		return types.ErrNameDoesNotExist(types.DefaultCodespace).Result()
	}

	err := keeper.DeleteShareHolder(ctx, msg.Validator, msg.Delegator)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{}
}
