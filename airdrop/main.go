package main

import (
	"database/sql"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/airdrop/event"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/airdrop/transfer"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/airdrop/util"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//todo:2 using the env and mount

//default air drop pool address

//todo, add time limit, if a transaction is not success in 24 hours, we should cancel it.

func readConfigFile() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("can not get working path .. ", err)
		return nil, err
	}

	fmt.Println("current working path：", wd)

	data, err := ioutil.ReadFile("./airdrop/config.yaml")
	if err != nil {
		log.Fatalf("can not read file ：%v", err)
	}

	// parse config file.
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("can not decode file ：%v", err)
	}
	return &config, nil
	// using config file
}

func main() {
	config, err := readConfigFile()
	if err != nil {
		return
	}

	util.Address, _ = bfc_types.NewAddressFromHex(config.Node.Address)
	util.M1Mnemonic = config.Node.M1Mnemonic
	handler := util.InitDB()
	ticker := time.NewTicker(1 * 60 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				event.QueryEvent(handler, config.Node.Rpc, event.CrossChainAssetsEvent)
				event.QueryTransaction(handler, config.Node.Rpc, config.Node.OwltoFromAddress)
				transfer.TransferBFC(handler, config.Node.Amount, config.Node.Rpc)
				OutPutCurrentAirDropStatus(handler)
				transfer.GetAddressBalance(config.Node.Address, config.Node.Rpc)
			}
		}
	}()

	select {}
}

func OutPutCurrentAirDropStatus(handler *sql.DB) {
	//todo, check current db status and output to console. [both todo airdrop count, and airdrop success count]
	successCount, _ := util.GetAirdropCountByStatus(handler, true)
	failedCount, _ := util.GetAirdropCountByStatus(handler, false)

	fmt.Printf("airdrop success count is %d,  todo airdrop count is %d\n", successCount, failedCount)
	//or output the latest 5 transactions data content.
}
