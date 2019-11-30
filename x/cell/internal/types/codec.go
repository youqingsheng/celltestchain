package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetShareHolder{}, "cell/SetShareHolder", nil)
	cdc.RegisterConcrete(MsgDeleteShareHolder{}, "cell/DeleteShareHolder", nil)
	//cdc.RegisterConcrete(MsgDeleteName{}, "nameservice/DeleteName", nil)

	cdc.RegisterConcrete(Share{}, "cell/Share", nil)

	cdc.RegisterConcrete(ShareHolder{}, "cell/ShareHolder", nil)
}
