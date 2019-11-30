package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ShareHolder struct {
	Address sdk.AccAddress `json:"address"`
	Rate    sdk.Dec        `json:"rate"`
}

func NewShareHolder(address sdk.AccAddress, rate sdk.Dec) ShareHolder {
	return ShareHolder{
		Address: address,
		Rate:    rate,
	}
}

type Share struct {
	Owner        sdk.ValAddress `json:"owner"`
	ShareHolders []ShareHolder  `json:"shareholders"`
}

func NewShare(owner sdk.ValAddress) Share {
	return Share{
		Owner: owner,
	}
}

// implement fmt.Stringer
func (s Share) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Shareholders: %s`, s.Owner, s.ShareHolders))
}
