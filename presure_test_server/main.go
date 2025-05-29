package main

import (
	"context"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/account"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/types"
	"github.com/gin-gonic/gin"
	"math/big"
)

func main() {
	println("hello")
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET(
		"/ping", func(c *gin.Context) {
			CallClient_onChain_PayBFC()
			//输出json结果给调用方
			c.JSON(
				200, gin.H{
					"message": "pong",
				},
			)
		},
	)
	r.Run(":13344") // listen and serve on 0.0.0.0:8080
}

func generateRandomRPCRequest() {
	//client.()
}

type BfcAddress = bfc_types.BfcAddress

func CallClient_onChain_PayBFC() {
	//cli, _ := client.Dial(types.TestnetRpcUrl)
	cli, _ := client.Dial(types.TestnetRpcUrl)

	signer, _ := account.NewAccountWithMnemonic(M1Mnemonic)

	recipient := Address

	coins, err := cli.GetCoins(context.Background(), *Address, nil, nil, 10)
	if err != nil {
		println("err = ", err.Error())
	}

	//cli.SplitCoin()

	amount := BFC(0.001).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 0, 0)

	txn, err := cli.PayBFC(
		context.Background(),
		*Address,
		pickedCoins.CoinIds(),
		[]BfcAddress{*recipient},
		[]types.SafeBfcBigInt[uint64]{
			types.NewSafeBfcBigInt(amount),
		},
		types.NewSafeBfcBigInt(gasBudget),
	)
	// sign and send
	signature, err := signer.SignSecureWithoutEncode(txn.TxBytes, bfc_types.DefaultIntent())
	options := types.BfcTransactionBlockResponseOptions{
		ShowEffects: true,
	}
	resp, err := cli.ExecuteTransactionBlock(
		context.TODO(), txn.TxBytes, []any{signature}, &options,
		types.TxnRequestTypeWaitForLocalExecution,
	)

	println("resp = ", resp)
}

func CallClient_onChain_SplitBFC() {
	cli, _ := client.Dial(types.TestnetRpcUrl)

	signer, _ := account.NewAccountWithMnemonic(M1Mnemonic)

	//recipient := Address

	coins, err := cli.GetCoins(context.Background(), *Address, nil, nil, 10)
	if err != nil {
		println("err = ", err)
	}

	//cli.SplitCoin()

	splitCoin := BFC(0.001).Uint64()
	amount := BFC(0.001).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 0, 0)

	txn, err := cli.SplitCoin(
		context.Background(),
		*Address,
		pickedCoins.Coins[0].CoinObjectId,
		[]types.SafeBfcBigInt[uint64]{
			types.NewSafeBfcBigInt(splitCoin),
		},
		nil,
		types.NewSafeBfcBigInt(gasBudget),
	)
	// sign and send
	signature, err := signer.SignSecureWithoutEncode(txn.TxBytes, bfc_types.DefaultIntent())
	options := types.BfcTransactionBlockResponseOptions{
		ShowEffects: true,
	}
	resp, err := cli.ExecuteTransactionBlock(
		context.TODO(), txn.TxBytes, []any{signature}, &options,
		types.TxnRequestTypeWaitForLocalExecution,
	)

	println("resp = ", resp)
}
