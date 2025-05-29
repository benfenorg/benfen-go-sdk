package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/rollback_transaction_replay/common"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/types"
	jsoniter "github.com/json-iterator/go"
)

var basePath = "/Users/wubin/workspace/obc/target/debug/bfc"
var epoch = "10"
var checkPoint = "94"
var db_handler *sql.DB = nil
var cli, _ = client.Dial(types.DevIndexRpc)

func main() {

	//get checkpoints from epoch...
	//for...
	db_handler = common.InitDB()
	// 1. save data into database
	first, end, err := common.GetAllCheckPoints(cli, "3")
	end = 1043
	if err != nil {
		fmt.Printf("GetAllCheckPoints failed, caused by %s", err.Error())
		return
	}
	for i := first; i <= end; i++ {
		saveTransactionDataInCheckPoint(i)
	}

	// 2. replay all not system transaction
	//saveTransactionDataInCheckPoint(273)
	//for i := first; i <= end; i++ {
	//	ReplayTransaction(db_handler, i)
	//}
}

func saveTransactionDataInCheckPoint(check_point int64) {
	transactions, err := common.GetTransactionFromCheckpoints(cli, strconv.FormatInt(check_point, 10))
	if err != nil {
		panic(err)
	}
	for _, digest := range transactions {
		if common.IsSystemContract(cli, digest, context.Background()) {
			fmt.Printf("digest %s is system contract, not need to replay", digest)
			common.InsertData(db_handler, check_point, digest, "skip")
			continue
		}

		txHash, txSign := common.GetTransactionByDigest(digest, cli)
		if len(txHash) == 0 {
			continue
		}

		rawTransaction := common.DecodeRawTransaction(strings.Trim(txHash, "\""), basePath)
		input := &types.TxSign{
			Hash: txSign,
		}
		tx, err := jsoniter.MarshalToString(input)
		if err != nil {
			fmt.Printf("MarshalToString failed caused by %s\n", err.Error())
			continue
		}
		common.SaveData(db_handler, digest, rawTransaction, tx, check_point)
	}
}
