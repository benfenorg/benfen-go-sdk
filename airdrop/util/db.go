package util

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	// 连接到SQLite数据库
	db, err := sql.Open("sqlite3", "airdrop.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	db.SetMaxOpenConns(5)

	// 创建表
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS airdrop (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		address TEXT NOT NULL,
		txDigest TEXT NOT NULL,
		isAirdropped BOOLEAN NOT NULL,
    	UNIQUE (txDigest)
	)`,
	)
	if err != nil {
		fmt.Println("error is ", err.Error())
		return nil
	}

	return db
}

func UpdateData(db *sql.DB, digest string) {
	_, err := db.Exec(
		"UPDATE airdrop set isAirdropped = true  where txDigest = ?",
		digest,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func SaveData(db *sql.DB, address, digest string) {
	_, err := db.Exec(
		"INSERT INTO airdrop (address, txDigest, isAirdropped) VALUES (?, ?, ?)",
		address,
		digest,
		false,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

type Data struct {
	ID           int    `db:"id"`
	Address      string `db:"address"`
	TxDigest     string `db:"txDigest"`
	IsAirdropped bool   `db:"isAirdropped"`
}

func QueryTransactions(db *sql.DB) ([]*Data, error) {
	// 查询交易数据
	rows, err := db.Query(
		"SELECT id, address, txDigest, isAirdropped FROM airdrop WHERE isAirdropped = false",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*Data
	for rows.Next() {
		d := &Data{}
		err := rows.Scan(
			&d.ID,
			&d.Address,
			&d.TxDigest,
			&d.IsAirdropped,
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

func GetAirdropCountByStatus(db *sql.DB, status bool) (int64, error) {
	rows, err := db.Query(
		"SELECT COUNT(*)  FROM airdrop WHERE isAirdropped = ?",
		status,
	)
	if err != nil {
		return 0, err
	}
	total := int64(0)
	for rows.Next() {
		err := rows.Scan(
			&total,
		)
		if err != nil {
			fmt.Println("GetAirdropCountByStatus failed, error: ", err)
			continue
		}
	}
	defer rows.Close()
	return total, nil
}
