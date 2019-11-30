package cell

import (
	"github.com/cell-network/cellchain/x/cell/internal/keeper"
	"github.com/cell-network/cellchain/x/cell/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper               = keeper.NewKeeper
	NewQuerier              = keeper.NewQuerier
	NewMsgDeleteShareHolder = types.NewMsgDeleteShareHolder
	NewMsgSetShareHolder    = types.NewMsgSetShareHolder
	NewShareHolder          = types.NewShareHolder
	ModuleCdc               = types.ModuleCdc
	RegisterCodec           = types.RegisterCodec
)

type (
	Keeper               = keeper.Keeper
	MsgSetShareHolder    = types.MsgSetShareHolder
	MsgDeleteShareHolder = types.MsgDeleteShareHolder
	ShareHolder          = types.ShareHolder
	Share                = types.Share
)
