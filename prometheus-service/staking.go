package main

import (
	"context"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/types"
	"strconv"
	"time"
)

var whiteList = map[string]string{
	"00000000000000000000000000000000000000000000000000000000000000c8::busd::BUSD": "0xc8::busd::BUSD",
	"00000000000000000000000000000000000000000000000000000000000000c8::bjpy::BJPY": "0xc8::bjpy::BJPY",
}

func getStakingJob() {
	getStakingInfo()

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

func getStakingInfo() {
	chain, _ := client.Dial(types.TestnetRpcUrl)
	resp, err := chain.GetLatestSuiSystemState(
		context.Background(),
	)

	if err != nil {
		println("get staking info failed", err)
	}
	for _, node := range resp.ActiveValidators {

		hexString := convertBfcAddressToHexAddress(node.StablePools.StablePoolsId.Sid)

		objID, err := bfc_types.NewAddressFromHex(
			hexString,
		)

		resp_stable_pools, err := chain.GetStablePools(
			context.Background(),
			*objID,
		)

		if err != nil {
			println("get stablePools info failed", err)
		}

		//println("the result is ", resp_stable_pools)

		arraydata, err := ConvertPoolObject(resp_stable_pools)
		if err != nil {
			println("convert pool object failed", err)
		}
		for index, poolobj := range arraydata {
			println("the index ", index, " poolobj is ", poolobj.ObjectID)
			PutStablePoolsInfoIntoMetric(poolobj.ObjectID, poolobj.ObjectID, chain)
		}
	}
}

func PutStablePoolsInfoIntoMetric(
	poolIdString string,
	stablePoolsId string, chain *client.Client,
) {
	hexAddress := convertBfcAddressToHexAddress(poolIdString)
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

	//println("the Content is ", got)

	stablePoolDetailInfo, err := ConvertToStablePoolDetailInfo(got)
	var stablePoolName string

	if v, ok := whiteList[stablePoolDetailInfo.Data.Content.Fields.Name]; ok {
		stablePoolName = v
	} else {
		return
	}

	stableBalance, _ := strconv.ParseFloat(stablePoolDetailInfo.Data.Content.Fields.Value.Fields.Stable_balance, 64)
	rewardsPool, _ := strconv.ParseFloat(stablePoolDetailInfo.Data.Content.Fields.Value.Fields.Rewards_pool, 64)
	bfcSystemMonitor.StakingInfo.WithLabelValues(
		"staking",
		stablePoolsId,
		poolIdString,
		stablePoolName,
		"stable_balance",
	).Set(stableBalance)

	bfcSystemMonitor.StakingInfo.WithLabelValues(
		"staking",
		stablePoolsId,
		poolIdString,
		stablePoolName,
		"rewards_pool",
	).Set(rewardsPool)
}

func GetSuiSystemStakingData() {
	chain, _ := client.Dial(types.TestnetRpcUrl)

	resp, err := chain.GetLatestSuiSystemState(
		context.Background(),
	)

	if err != nil {
		println("get dao info failed", err)
	}

	for _, node := range resp.ActiveValidators {
		stakeingPoolBalance := node.StakingPoolBfcBalance
		println("stakeingPoolBfcBalance", stakeingPoolBalance.Uint64())
		stablePools := node.StablePools
		//stablePoolsId := node.StablePools.StablePoolsId.Sid
		println("the stable pools id is ", node.StablePools.StablePoolsId.Sid)
		println("the stable pools size is ", stablePools.Size.Uint64())

		hexString := convertBfcAddressToHexAddress(node.StablePools.StablePoolsId.Sid)

		objID, err := bfc_types.NewAddressFromHex(
			hexString,
		)

		resp_stable_pools, err := chain.GetStablePools(
			context.Background(),
			*objID,
		)

		if err != nil {
			println("get stablePools info failed", err)
		}

		//println("the result is ", resp_stable_pools)

		arraydata, err := ConvertPoolObject(resp_stable_pools)
		if err != nil {
			println("convert pool object failed", err)
		}
		for index, poolobj := range arraydata {
			println("the index ", index, " poolobj is ", poolobj.ObjectID)
			GetStablePoolsInfo(poolobj.ObjectID, node.StablePools.StablePoolsId.Sid, chain)
		}
	}

	time.Sleep(5 * time.Second)
	//common.PrintJson(resp)
}
