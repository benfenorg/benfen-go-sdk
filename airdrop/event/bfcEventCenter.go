package event

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/airdrop/util"
	"github.com/benfenorg/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/types"
	"strings"
)

var CrossChainAssetsEvent = "0x02079953f9340da8d0863b62853b7443214ec5f48f0e7931b9b50222580ede5f::event::UnlockedEvent"
var ValidatorInfoEvent = "0x3::validator_set::ValidatorEpochInfoEventV2"

func QueryValidatorInfoEvent(rpcUrl string, eventType string) {
	var cli, _ = client.Dial(rpcUrl)
	limit := uint(100)

	data, err := cli.QueryEvents(
		context.Background(),
		types.EventFilter{
			MoveEventType: &eventType,
		},
		nil,
		&limit,
		true,
	)
	if err != nil || data == nil || len(data.Data) == 0 {
		fmt.Println("err = or data is nil", err)
		return
	}

	for _, v := range data.Data {
		input := v.ParsedJson.(map[string]interface{})
		stablePoolExchangeData := input["stable_pool_token_exchange_rate"]

		validatorAddress := input["validator_address"]
		fmt.Println("validator_address: ", validatorAddress)
		fmt.Println("data : ", stablePoolExchangeData)
	}
	//util.PrintJson(data)
}

func QueryEvent(handler *sql.DB, rpcUrl string, eventType string) {
	var cli, _ = client.Dial(rpcUrl)
	limit := uint(100)

	data, err := cli.QueryEvents(
		context.Background(),
		types.EventFilter{
			MoveEventType: &eventType,
		},
		nil,
		&limit,
		true,
	)
	if err != nil || data == nil || len(data.Data) == 0 {
		fmt.Println("err = or data is nil", err)
		return
	}

	for _, v := range data.Data {
		input := v.ParsedJson.(map[string]interface{})

		println("to_address:", input["recipient"].(string), ",dist: ", v.Id.TxDigest.String())
		txTime := v.TimestampMs.Uint64()
		currentTime := util.GetTimeStamp() * 1000
		if currentTime-txTime > 3*24*60*60*1000 {
			fmt.Println("skip this trx,.... time too long.....3 days ago.")
			continue
		}
		util.SaveData(handler, input["recipient"].(string), v.Id.TxDigest.String())
	}
	//util.PrintJson(data)
	util.PrintJson("save ok.")
}

func QueryTransaction(handler *sql.DB, rpcUrl string, owltoFromAddress string) {
	var cli, _ = client.Dial(rpcUrl)
	limit := uint(100)
	newAddress, _ := bfc_types.NewAddressFromHex(owltoFromAddress)

	data, err := cli.QueryTransactionBlocks(
		context.Background(),
		types.BfcTransactionBlockResponseQuery{
			Filter: &types.TransactionFilter{
				FromAddress: newAddress,
			},
		},
		nil,
		&limit,
		true,
	)
	if err != nil || data == nil || len(data.Data) == 0 {
		fmt.Println("err = or data is nil", err)
		return
	}
	for _, v := range data.Data {
		resp, err := cli.GetTransactionBlock(
			context.Background(), v.Digest, types.BfcTransactionBlockResponseOptions{
				ShowInput:          true,
				ShowEffects:        true,
				ShowObjectChanges:  true,
				ShowBalanceChanges: true,
				ShowEvents:         true,
			},
		)
		if err != nil || resp == nil {
			fmt.Println("GetTransactionBlock data is nil", err)
			continue
		}
		in := resp.Transaction.Data.Data.V1.Transaction.Data.ProgrammableTransaction.Inputs
		change := resp.BalanceChanges
		found := false
		for _, balanceChange := range change {
			if strings.Contains(balanceChange.CoinType, "::busd::BUSD") {
				found = true
			}
		}
		if len(in) >= 3 && found {
			input := in[len(in)-1]
			toAddress := input.(map[string]interface{})["value"].(string)
			println("owlto to_address:", toAddress, ",dist: ", v.Digest.String())
			txTime := resp.TimestampMs.Uint64()
			currentTime := util.GetTimeStamp() * 1000
			if currentTime-txTime > 3*24*60*60*1000 {
				fmt.Println("owlto skip this trx,.... time too long.....3 days ago.")
				continue
			}
			util.SaveData(handler, toAddress, v.Digest.String())
		}
	}
	util.PrintJson("save ok.")
}
