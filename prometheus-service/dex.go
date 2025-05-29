package main

import (
	"context"
	"github.com/benfenorg/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/types"
	"strconv"
	"time"
)

const BusdLong = "0x1537ad845c190d7363b8e558100a5d782c2d4ce6e2a39d0f1b56db639a3ad3fa"
const UsdcLong = "0xa660debbc60f9c131c82728b8fae18335622533bdb0206f83558f0d4c95dd448"

func getDexInfo() {
	getDexInfoByAddress(BusdLong, "busdlong")
	getDexInfoByAddress(UsdcLong, "usdclong")
}

func getDexJob() {
	getDexInfo()

	ticker1 := time.NewTicker(60 * 60 * time.Second)
	defer ticker1.Stop()
	go func(t *time.Ticker) {
		for {
			// 每5秒中从chan t.C 中读取一次
			<-t.C
			getDexInfo()
		}
	}(ticker1)
}

func getDexInfoByAddress(address, typeName string) {
	cli, _ := client.Dial(types.TestnetRpcUrl)
	objId, _ := bfc_types.NewAddressFromHex(address)
	obj, err := cli.GetObject(
		context.Background(), *objId, &types.BfcObjectDataOptions{
			ShowType:    true,
			ShowContent: true,
		},
	)
	if err != nil {
		println("get dex info failed, caused by %s", err.Error())
	}

	var fields any
	if obj != nil &&
		obj.Data != nil &&
		obj.Data.Content != nil &&
		obj.Data.Content.Data.MoveObject != nil {
		fields = obj.Data.Content.Data.MoveObject.Fields
	} else {
		println("can not GetObject")
	}

	input := fields.(map[string]interface{})
	coinA, _ := strconv.ParseFloat(input["coin_a"].(string), 64)
	coinB, _ := strconv.ParseFloat(input["coin_b"].(string), 64)
	currentSqrtPrice, _ := strconv.ParseFloat(input["current_sqrt_price"].(string), 64)

	bfcSystemMonitor.DexInfo.WithLabelValues(typeName, "coinABalance").Set(coinA)
	bfcSystemMonitor.DexInfo.WithLabelValues(typeName, "CoinBBalance").Set(coinB)
	bfcSystemMonitor.DexInfo.WithLabelValues(typeName, "CurrentSqrtPrice").Set(currentSqrtPrice)
}
