package main

import (
	"context"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/types"
	"testing"
)

func Test_getNFT(t *testing.T) {
	chain, _ := client.Dial(types.TestnetRpcUrl)
	data := "BFC8f57efbbf10ea5af8429ab622be81de374b4e5edd0a71a0afe8a597e2467ac6a9d6b"
	hexAddress := convertBfcAddressToHexAddress(data)
	objectID, _ := bfc_types.NewAddressFromHex(hexAddress)
	got, err := chain.GetObjectString(
		context.Background(), *objectID, &types.BfcObjectDataOptions{
			ShowType:                true,
			ShowOwner:               true,
			ShowContent:             true,
			ShowDisplay:             true,
			ShowBcs:                 true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	)

	if err != nil {
		println("get stablePools detail info failed", err)
	}

	println("the Content is ", got)

	NftInfo, err := ConvertToNftDetailInfo(got)

	println("the name is ", NftInfo.Data.Content.Fields.Name)
	println("the BfcBalance is ", NftInfo.Data.Content.Fields.Value.Fields.BfcBalance)
	println("the Period is ", NftInfo.Data.Content.Fields.Value.Fields.Period)
	println("the TotalPower is ", NftInfo.Data.Content.Fields.Value.Fields.TotalPower)

}
