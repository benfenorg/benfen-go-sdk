package main

import (
	"context"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/rollback_transaction_replay/common"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/types"
	"time"
)

func getDaoJob() {
	//todo using time to update the dao info
	getDaoObjectInfo()
	ticker1 := time.NewTicker(5 * 60 * time.Second)
	defer ticker1.Stop()
	go func(t *time.Ticker) {
		for {
			// 每5秒中从chan t.C 中读取一次
			<-t.C
			getDaoObjectInfo()
		}
	}(ticker1)
}

func getDaoObjectInfo() {
	chain, _ := client.Dial(types.TestnetRpcUrl)

	resp, err := chain.GetInnerDaoInfoObject(
		context.Background(),
	)

	bfcSystemMonitor.DaoVotingPeriod.Set(float64(resp.Config.VotingPeriod / 1000))
	bfcSystemMonitor.DaoVotingQuorum.Set(float64(resp.Config.VotingQuorum))
	bfcSystemMonitor.DaoVotingActionDelay.Set(float64(resp.Config.VotingActionDelay / 1000))
	bfcSystemMonitor.DaoVotingDelay.Set(float64(resp.Config.VotingDelay / 1000))

	if err != nil {
		println("get dao info failed", err)
	}

	common.PrintJson(resp)
}
