package main

import (
	"context"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/account"
	"github.com/benfenorg/benfen-go-sdk/airdrop/util"
	"github.com/benfenorg/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/types"
	"math/big"
)

func TransferBFCOnChain(data string, airdropAmount float64, rpcUrl string) {
	hexAddress := util.ConvertBfcAddressToHexAddress(data)
	cli, _ := client.Dial(rpcUrl)
	sender, err := bfc_types.NewAddressFromHex(util.Address.String())

	signer, _ := account.NewAccountWithMnemonic(util.M1Mnemonic)

	recipient, _ := bfc_types.NewAddressFromHex(hexAddress)

	coins, err := cli.GetCoins(context.Background(), *util.Address, nil, nil, 10)
	if err != nil || coins.Data == nil || len(coins.Data) == 0 {
		println("err = or data is nil", err.Error())
	}

	amount := util.BFC(airdropAmount).Uint64()
	gasBudget := util.BFC(0.001).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 0, 0)

	txn, err := cli.TransferBFC(
		context.Background(), *sender, *recipient, pickedCoins.Coins[0].CoinObjectId,
		types.NewSafeBfcBigInt(amount),
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

	if err != nil {
		fmt.Printf("err is :%s \n", err.Error())
		return
	}
	println("resp =", resp.Digest.String())
}
