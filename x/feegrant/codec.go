package feegrant

import (
	"github.com/weijun-sh/cosmos-sdk/codec"
	"github.com/weijun-sh/cosmos-sdk/codec/legacy"
	"github.com/weijun-sh/cosmos-sdk/codec/types"
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	"github.com/weijun-sh/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/feegrant interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgGrantAllowance{}, "cosmos-sdk/MsgGrantAllowance", nil)
	cdc.RegisterConcrete(&MsgRevokeAllowance{}, "cosmos-sdk/MsgRevokeAllowance", nil)

	cdc.RegisterInterface((*FeeAllowanceI)(nil), nil)
	cdc.RegisterConcrete(&BasicAllowance{}, "cosmos-sdk/BasicAllowance", nil)
	cdc.RegisterConcrete(&PeriodicAllowance{}, "cosmos-sdk/PeriodicAllowance", nil)
	cdc.RegisterConcrete(&AllowedMsgAllowance{}, "cosmos-sdk/AllowedMsgAllowance", nil)
}

// RegisterInterfaces registers the interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGrantAllowance{},
		&MsgRevokeAllowance{},
	)

	registry.RegisterInterface(
		"cosmos.feegrant.v1beta1.FeeAllowanceI",
		(*FeeAllowanceI)(nil),
		&BasicAllowance{},
		&PeriodicAllowance{},
		&AllowedMsgAllowance{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/feegrant module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/feegrant and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the global Amino codec so that this can later be
	// used to properly serialize x/authz MsgExec instances
	RegisterLegacyAminoCodec(legacy.Cdc)
}
