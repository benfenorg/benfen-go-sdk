package main

import (
	"context"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/types"
	"testing"
)

func Test_get_system(t *testing.T) {
	GetSuiSystemStakingData()
}

func Test_Convert(t *testing.T) {
	data := "BFC3ec9d6c82344f251dd02f57ca8305b5654f3a249719b246504d11bdf3d324038fc89"
	result := convertBfcAddressToHexAddress(data)

	println(result)
}

func Test_getStablePools(t *testing.T) {
	chain, _ := client.Dial(types.TestnetRpcUrl)

	//this is the stalbe pool address
	data := "BFC3ec9d6c82344f251dd02f57ca8305b5654f3a249719b246504d11bdf3d324038fc89"
	hexAddress := convertBfcAddressToHexAddress(data)
	objectID, _ := bfc_types.NewAddressFromHex(hexAddress)
	resp, err := chain.GetStablePools(
		context.Background(),
		*objectID,
	)

	if err != nil {
		println("get stablePools info failed", err)
	}

	println("the result is ", resp)

	arraydata, err := ConvertPoolObject(resp)
	if err != nil {
		println("convert pool object failed", err)
	}
	for _, poolobj := range arraydata {
		println("the result is ", poolobj.ObjectID)
	}

}

func Test_getStablePoolsInfo(t *testing.T) {
	chain, _ := client.Dial(types.TestnetRpcUrl)
	data := "BFC00aa563ae5af52d7d797a06e27ab35598b9c27c27454380e6141ceda65bb077ffe0f"
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

	stablePoolDetailInfo, err := ConvertToStablePoolDetailInfo(got)

	println("the name is ", stablePoolDetailInfo.Data.Content.Fields.Name)
	println("the Stable_balance is ", stablePoolDetailInfo.Data.Content.Fields.Value.Fields.Stable_balance)
	println("the Rewards_pool is ", stablePoolDetailInfo.Data.Content.Fields.Value.Fields.Rewards_pool)

}
