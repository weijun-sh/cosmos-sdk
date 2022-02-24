package params

import (
	"github.com/weijun-sh/cosmos-sdk/client"
	"github.com/weijun-sh/cosmos-sdk/codec"
	"github.com/weijun-sh/cosmos-sdk/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}
