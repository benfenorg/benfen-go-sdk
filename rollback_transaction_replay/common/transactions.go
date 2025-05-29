package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/types"
)

func GetTransactionByDigest(digest string, cli *client.Client) (string, []string) {
	context := context.Background()
	//require.NoError(t, err)
	options := types.BfcTransactionBlockResponseOptions{
		ShowInput:    true,
		ShowRawInput: true,
		ShowEffects:  true,
		ShowEvents:   true,
	}

	d, err := bfc_types.NewDigest(digest)
	resp, err := cli.GetTransactionBlock(
		context,
		*d,
		options,
	)
	if err != nil {
		println("err = ", err.Error())
		return "", nil
	}

	PrintJson(resp)

	println("resp.RawTransaction ")
	body, _ := json.Marshal(resp.RawTransaction)
	var str bytes.Buffer
	_ = json.Indent(&str, body, "", "    ")
	fmt.Println(str.String())

	return str.String(), resp.Transaction.TxSignatures
}
