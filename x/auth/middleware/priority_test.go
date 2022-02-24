package middleware_test

import (
	cryptotypes "github.com/weijun-sh/cosmos-sdk/crypto/types"
	"github.com/weijun-sh/cosmos-sdk/testutil/testdata"
	sdk "github.com/weijun-sh/cosmos-sdk/types"
	"github.com/weijun-sh/cosmos-sdk/types/tx"
	"github.com/weijun-sh/cosmos-sdk/x/auth/middleware"
)

func (s *MWTestSuite) TestPriority() {
	ctx := s.SetupTest(true) // setup
	txBuilder := s.clientCtx.TxConfig.NewTxBuilder()

	txHandler := middleware.ComposeMiddlewares(noopTxHandler, middleware.TxPriorityMiddleware)

	// keys and addresses
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	// msg and signatures
	msg := testdata.NewTestMsg(addr1)
	atomCoin := sdk.NewCoin("atom", sdk.NewInt(150))
	apeCoin := sdk.NewInt64Coin("ape", 1500000)
	feeAmount := sdk.NewCoins(apeCoin, atomCoin)
	gasLimit := testdata.NewTestGasLimit()
	s.Require().NoError(txBuilder.SetMsgs(msg))
	txBuilder.SetFeeAmount(feeAmount)
	txBuilder.SetGasLimit(gasLimit)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	testTx, _, err := s.createTestTx(txBuilder, privs, accNums, accSeqs, ctx.ChainID())
	s.Require().NoError(err)

	// txHandler errors with insufficient fees
	_, checkTxRes, err := txHandler.CheckTx(sdk.WrapSDKContext(ctx), tx.Request{Tx: testTx}, tx.RequestCheckTx{})
	s.Require().NoError(err, "Middleware should not have errored on too low fee for local gasPrice")
	s.Require().Equal(atomCoin.Amount.Int64(), checkTxRes.Priority, "priority should be atom amount")
}
