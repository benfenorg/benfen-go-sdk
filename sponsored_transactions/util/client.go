package util

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/account"
	"github.com/benfenorg/benfen-go-sdk/airdrop/util"
	"github.com/benfenorg/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/types"
	"github.com/fardream/go-bcs/bcs"
)

func TransferObjectAllowSponsor(
	senderPublicKey, sponsorPublicKey, senderMnemonic, sponsorMnemonic string,
	gasPrice uint64, rpcUrl string,
) {
	sender, err := bfc_types.NewAddressFromHex(senderPublicKey)
	sender2, _ := bfc_types.NewAddressFromHex(sponsorPublicKey)
	signer, _ := account.NewAccountWithMnemonic(senderMnemonic)
	signer2, _ := account.NewAccountWithMnemonic(sponsorMnemonic)

	recipient := sender
	gasBudget := util.BFC(0.01).Uint64()

	cli, _ := client.Dial(rpcUrl)

	coins, err := cli.GetCoins(context.Background(), *sender, nil, nil, uint(10))
	coin := coins.Data[0]

	gases, err := cli.GetCoins(context.Background(), *sender2, nil, nil, uint(10))
	gas := gases.Data[0]
	// gasPrice, err := cli.GetReferenceGasPrice(context.Background())

	// build with BCS
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	err = ptb.TransferObject(*recipient, []*bfc_types.ObjectRef{coin.Reference()})
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammableAllowSponsor(
		*sender, []*bfc_types.ObjectRef{
			gas.Reference(),
		},
		pt, gasBudget, gasPrice, *sender2,
	)
	txBytesBCS, err := bcs.Marshal(tx)

	fmt.Println(base64.StdEncoding.EncodeToString(txBytesBCS))

	signature, err := signer.SignSecureWithoutEncode(txBytesBCS, bfc_types.DefaultIntent())

	signature2, err := signer2.SignSecureWithoutEncode(txBytesBCS, bfc_types.DefaultIntent())

	options := types.BfcTransactionBlockResponseOptions{
		ShowEffects: true,
	}
	resp, err := cli.ExecuteTransactionBlock(
		context.TODO(), txBytesBCS, []any{signature, signature2}, &options,
		types.TxnRequestTypeWaitForLocalExecution,
	)

	if err != nil {
		fmt.Printf("err is :%s \n", err.Error())
		return
	}
	fmt.Printf("resp digest is : %s\n", resp.Digest)
}
