package types

import (
	"github.com/weijun-sh/cosmos-sdk/codec"
	"github.com/weijun-sh/cosmos-sdk/codec/legacy"
	"github.com/weijun-sh/cosmos-sdk/codec/types"
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	"github.com/weijun-sh/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers concrete types on LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUnjail{}, "cosmos-sdk/MsgUnjail", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnjail{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(legacy.Cdc)
}
