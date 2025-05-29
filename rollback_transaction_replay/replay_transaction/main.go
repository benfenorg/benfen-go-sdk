package main

import (
	"database/sql"
	"fmt"

	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/rollback_transaction_replay/common"
	"github.com/benfenorg/benfen-go-sdk/rollback_transaction_replay/model"
	"github.com/benfenorg/benfen-go-sdk/types"
)

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

	// 2. replay all not system transaction
	//saveTransactionDataInCheckPoint(273)
	var allData []*model.Data

	for i := first; i <= end; i++ {
		data, err := common.QueryTransactions(db_handler, i)
		if err != nil {
			fmt.Printf("QueryTransactions failed, caused by %s", err.Error())
			return
		}
		allData = append(allData, data...)
		// common.ReplayTransaction(cli, db_handler, i)
	}
	for _, data := range allData {
		err := common.ReplayAndInsertTransactionData(cli, db_handler, data)
		if err != nil {
			fmt.Printf("InsertTransactionData failed, caused by %s", err.Error())
			return
		}
		// common.ReplayTransaction(cli, db_handler, data.Epoch)
	}
}
