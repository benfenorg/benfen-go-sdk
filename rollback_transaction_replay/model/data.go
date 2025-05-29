package model

type Data struct {
	ID              int    `db:"id"`
	Epoch           int64  `db:"epoch"`
	Digest          string `db:"digest"`
	TransactionData string `db:"transaction_data"`
	Sign            string `db:"sign"`
}
