package types

import (
	"github.com/weijun-sh/cosmos-sdk/codec"
	cryptocodec "github.com/weijun-sh/cosmos-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
