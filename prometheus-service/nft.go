package main

import (
	"context"
	"github.com/benfenorg/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/types"
	"strconv"
	"time"
)

const NftAddress = "BFC8f57efbbf10ea5af8429ab622be81de374b4e5edd0a71a0afe8a597e2467ac6a9d6b"

func getNftJob() {
	getNftInfo()

	ticker1 := time.NewTicker(60 * 60 * time.Second)
	defer ticker1.Stop()
	go func(t *time.Ticker) {
		for {
			// 每5秒中从chan t.C 中读取一次
			<-t.C
			getNftInfo()
		}
	}(ticker1)
}

func getNftInfo() {
	cli, _ := client.Dial(types.TestnetRpcUrl)
	hexAddress := convertBfcAddressToHexAddress(NftAddress)
	objectID, _ := bfc_types.NewAddressFromHex(hexAddress)
	got, err := cli.GetObjectString(
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
		println("get nft detail info failed", err)
	}
	NftInfo, err := ConvertToNftDetailInfo(got)
	totalPower, _ := strconv.ParseFloat(NftInfo.Data.Content.Fields.Value.Fields.TotalPower, 64)
	bfcBalance, _ := strconv.ParseFloat(NftInfo.Data.Content.Fields.Value.Fields.BfcBalance, 64)
	period, _ := strconv.ParseFloat(NftInfo.Data.Content.Fields.Value.Fields.Period, 64)

	bfcSystemMonitor.NftInfo.WithLabelValues("NFT", "Period").Set(period)
	bfcSystemMonitor.NftInfo.WithLabelValues("NFT", "BfcBalance").Set(bfcBalance)
	bfcSystemMonitor.NftInfo.WithLabelValues("NFT", "TotalPower").Set(totalPower)
}
