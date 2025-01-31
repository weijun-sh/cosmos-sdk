package tmservice

import (
	"context"
	"github.com/weijun-sh/cosmos-sdk/client"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/tendermint/tendermint/rpc/coretypes"
)

func getBlock(ctx context.Context, clientCtx client.Context, height *int64) (*coretypes.ResultBlock, error) {
	// get the node
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	return node.Block(ctx, height)
}

func GetProtoBlock(ctx context.Context, clientCtx client.Context, height *int64) (tmproto.BlockID, *tmproto.Block, error) {
	block, err := getBlock(ctx, clientCtx, height)
	if err != nil {
		return tmproto.BlockID{}, nil, err
	}
	protoBlock, err := block.Block.ToProto()
	if err != nil {
		return tmproto.BlockID{}, nil, err
	}
	protoBlockId := block.BlockID.ToProto()

	return protoBlockId, protoBlock, nil
}
