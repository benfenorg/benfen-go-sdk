package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func index_main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "dbuser1"
		password = "obcobc"
		dbname   = "obcdb"
	)
	epoch := 173
	checkpoint := 3402402

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	deleteQuery := fmt.Sprintf("delete from address_stats where checkpoint > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from changed_objects where checkpoint_sequence_number > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from checkpoint_metrics where checkpoint > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from checkpoints where sequence_number > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from dao_proposals where checkpoint_sequence_number > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf(
		"delete from dao_votes where transaction_digest in (select transaction_digest from transactions where checkpoint_sequence_number > %d);",
		checkpoint,
	)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from epoch_stake_coins where epoch > %d;", epoch)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from epoch_stakes where epoch > %d;", epoch)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from epochs where epoch > %d;", epoch)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf(
		"delete from events where transaction_digest in (select transaction_digest from transactions where checkpoint_sequence_number > %d);",
		checkpoint,
	)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from input_objects where checkpoint_sequence_number > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf(
		"delete from mining_nft_liquidities where transaction_digest in (select transaction_digest from transactions where checkpoint_sequence_number > %d);",
		checkpoint,
	)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from move_calls where checkpoint_sequence_number > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from objects where checkpoint > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from objects_history where checkpoint > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from recipients where checkpoint_sequence_number > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from system_states where epoch > %d;", epoch)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from transactions where checkpoint_sequence_number > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from validators where epoch > %d;", epoch)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from mining_nfts where sequence_number > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	deleteQuery = fmt.Sprintf("delete from mining_nft_staking where sequence_number > %d;", checkpoint)
	execBySql(err, db, deleteQuery)
	execBySql(err, db, "REFRESH MATERIALIZED VIEW mining_nfts_view;")
	execBySql(err, db, "REFRESH MATERIALIZED VIEW mining_nft_staking_view;")
}

func execBySql(err error, db *sql.DB, deleteQuery string) {
	result, err := db.Exec(deleteQuery)
	if err != nil {
		log.Fatal("Error executing delete query: ", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Error getting rows affected: ", err)
	}
	fmt.Printf("Rows affected: %d\n", rowsAffected)
}
