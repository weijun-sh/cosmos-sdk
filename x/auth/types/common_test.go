package types_test

import (
	"github.com/weijun-sh/cosmos-sdk/simapp"
)

var (
	ecdc                  = simapp.MakeTestEncodingConfig()
	appCodec, legacyAmino = ecdc.Codec, ecdc.Amino
)
