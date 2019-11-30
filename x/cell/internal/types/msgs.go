package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetShareHolder defines a SetShareHolder message
type MsgSetShareHolder struct {
	Validator sdk.ValAddress `json:"validator"`
	Delegator sdk.AccAddress `json:"delegator"`
	Rate      sdk.Dec        `json:"rate"`
}

func NewMsgSetShareHolder(valAddr sdk.ValAddress, deleAddr sdk.AccAddress, rate sdk.Dec) MsgSetShareHolder {
	return MsgSetShareHolder{
		Validator: valAddr,
		Delegator: deleAddr,
		Rate:      rate,
	}
}

// Route should return the name of the module
func (msg MsgSetShareHolder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetShareHolder) Type() string { return "set_shareholder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetShareHolder) ValidateBasic() sdk.Error {
	if msg.Validator.Empty() {
		return sdk.ErrInvalidAddress(msg.Validator.String())
	}
	if msg.Delegator.Empty() {
		return sdk.ErrInvalidAddress(msg.Delegator.String())
	}
	if msg.Rate.GTE(sdk.NewDec(1)) || msg.Rate.LTE(sdk.NewDec(0)) {
		return sdk.ErrUnknownRequest("Rate can not great than 1 or less than 0!")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetShareHolder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetShareHolder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Validator.Bytes())}
}

// MsgDeleteShareHolder defines a DeleteShareHolder message
type MsgDeleteShareHolder struct {
	Validator sdk.ValAddress `json:"validator"`
	Delegator sdk.AccAddress `json:"delegator"`
}

// NewDeleteShareHolder is a constructor function for DeleteShareHolder
func NewMsgDeleteShareHolder(valAddr sdk.ValAddress, deleAddr sdk.AccAddress) MsgDeleteShareHolder {
	return MsgDeleteShareHolder{
		Validator: valAddr,
		Delegator: deleAddr,
	}
}

// Route should return the name of the module
func (msg MsgDeleteShareHolder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteShareHolder) Type() string { return "delete_shareholder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteShareHolder) ValidateBasic() sdk.Error {
	if msg.Validator.Empty() {
		return sdk.ErrInvalidAddress(msg.Validator.String())
	}
	if msg.Delegator.Empty() {
		return sdk.ErrInvalidAddress(msg.Delegator.String())
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteShareHolder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteShareHolder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Validator.Bytes())}
}
