package common

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/rollback_transaction_replay/model"
	"github.com/benfenorg/benfen-go-sdk/types"
	jsoniter "github.com/json-iterator/go"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	// 连接到SQLite数据库
	db, err := sql.Open("sqlite3", "rollback_data.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxOpenConns(5)

	// 创建表
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS rollbackdata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		epoch INTEGER NOT NULL,
		transaction_data TEXT NOT NULL,
		result_data VARCHAR(40) NOT NULL
	)`,
	)
	if err != nil {
		fmt.Println("error is ", err.Error())
		return nil
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS transactiondata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		epoch INTEGER NOT NULL,
		digest VARCHAR(200) NOT NULL,
		transaction_data TEXT NOT NULL,
		sign VARCHAR(200) NOT NULL
	)`,
	)
	if err != nil {
		fmt.Println("error is ", err.Error())
		return nil
	}
	return db
}

func DBSample() {

	db := InitDB()

	InsertData(db, 100, "BcfhsB7AxGjR9fnGyYcaEhIvsvs7Ox4gC8b", "success")

	GetAllData(db)

}

func SaveData(db *sql.DB, digest, transaction_raw, sign string, check_point int64) {
	_, err := db.Exec(
		"INSERT INTO transactiondata (epoch, digest, transaction_data, sign) VALUES (?, ?, ?, ?)",
		check_point,
		digest,
		transaction_raw,
		sign,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func QueryTransactions(db *sql.DB, checkpoint int64) ([]*model.Data, error) {
	// 查询交易数据
	rows, err := db.Query(
		"SELECT id, epoch, digest, transaction_data, sign FROM transactiondata WHERE epoch = ?",
		checkpoint,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*model.Data
	for rows.Next() {
		d := &model.Data{}
		err := rows.Scan(
			&d.ID,
			&d.Epoch,
			&d.Digest,
			&d.TransactionData,
			&d.Sign,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil

}

func ReplayAndInsertTransactionData(cli *client.Client, db *sql.DB, data *model.Data) error {
	var txsign = &types.TxSign{}
	err := jsoniter.UnmarshalFromString(data.Sign, txsign)
	if err != nil {
		fmt.Println("unmarshal from string failed, caused by %s", err.Error())
		return err
	}
	options := types.BfcTransactionBlockResponseOptions{
		ShowEffects: true,
	}
	resp, err := cli.ExecuteTransactionBlockStr(
		context.TODO(), data.TransactionData, txsign.Hash, &options,
		types.TxnRequestTypeWaitForLocalExecution,
	)
	if err != nil {
		fmt.Printf("replay digest failed caused by %s\n", err.Error())
		InsertData(db, data.Epoch, data.Digest, "failed")
	} else {
		fmt.Printf("replay digest success, digest id is %s\n", resp.Digest.String())
		InsertData(db, data.Epoch, data.Digest, "success")
	}
	return nil
}

func InsertData(db *sql.DB, epoch int64, transaction_id, status string) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tx.Rollback()
	// 插入数据
	_, err = tx.Exec(
		"INSERT INTO rollbackdata (epoch, transaction_data, result_data) VALUES (?, ?, ?)",
		epoch,
		transaction_id,
		status,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}
}

func GetAllData(db *sql.DB) {
	// 查询数据
	rows, err := db.Query("SELECT id,epoch, transaction_data, result_data FROM rollbackdata")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var epoch int
		var result_data string
		var transaction_data string
		err := rows.Scan(&id, &epoch, &transaction_data, &result_data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(
			"ID: %d, epoch:%d,  transaction: %s, result: %s\n", id,
			epoch,
			transaction_data, result_data,
		)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

func GetAllTransaction(db *sql.DB) {
	// 查询数据
	rows, err := db.Query("SELECT id, epoch, digest, transaction_data, sign  FROM transactiondata")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var transactionData string
		var sign string
		var epoch int64
		err := rows.Scan(&id, &epoch, &transactionData, &sign)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(
			"ID: %d, epoch:%d transaction: %s, sign: %s\n", id,
			epoch, transactionData, sign,
		)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

func ReplayTransaction(cli *client.Client, db *sql.DB, check_point int64) {
	// 查询数据
	rows, err := db.Query(
		"SELECT id, epoch, digest, transaction_data, sign  FROM transactiondata WHere epoch = $1",
		check_point,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var transactionData string
		var sign, digest string
		var epoch int64
		err := rows.Scan(&id, &epoch, &digest, &transactionData, &sign)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(
			"ID: %d, epoch:%d, digest:%s, transaction: %s, sign: %s\n", id,
			epoch, digest, transactionData, sign,
		)

		var txsign = &types.TxSign{}
		err = jsoniter.UnmarshalFromString(sign, txsign)
		if err != nil {
			fmt.Println("unmarshal from string failed, caused by %s", err.Error())
			continue
		}
		options := types.BfcTransactionBlockResponseOptions{
			ShowEffects: true,
		}
		resp, err := cli.ExecuteTransactionBlockStr(
			context.TODO(), transactionData, txsign.Hash, &options,
			types.TxnRequestTypeWaitForLocalExecution,
		)
		if err != nil {
			fmt.Printf("replay digest failed caused by %s\n", err.Error())
			InsertData(db, check_point, digest, "failed")
		} else {
			fmt.Printf("replay digest success, digest id is %s\n", resp.Digest.String())
			InsertData(db, check_point, digest, "success")
		}
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return
	}
}
