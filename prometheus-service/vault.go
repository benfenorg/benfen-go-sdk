package main

import (
	"context"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/client"
	"time"
)

func getVaultInfo() {
	rpc := client.GetTestBenfenRpc("")
	info, _ := rpc.VaultInfo(context.Background(), ChainName, "0xc8::busd::BUSD")
	// assert.Equal(t, info.BasePoint, uint64(1000000000000))
	fmt.Printf("a balance: %d\n", info.CoinABalance)
	fmt.Printf("b balance: %d\n", info.CoinBBalance)
	ans, _ := client.SqrtPrice2Price(info.LastSqrtPrice.String())
	bfcSystemMonitor.VaultInfo.WithLabelValues("busd", "coinABalance").Set(float64(info.CoinABalance))
	bfcSystemMonitor.VaultInfo.WithLabelValues("busd", "CoinBBalance").Set(float64(info.CoinBBalance))
	bfcSystemMonitor.VaultInfo.WithLabelValues("busd", "price").Set(ans)
}

func getVaultJob() {
	//todo using time to update the dao info
	getVaultInfo()
	ticker1 := time.NewTicker(60 * 60 * time.Second)
	defer ticker1.Stop()
	go func(t *time.Ticker) {
		for {
			// 每5秒中从chan t.C 中读取一次
			<-t.C
			getVaultInfo()
		}
	}(ticker1)
}
