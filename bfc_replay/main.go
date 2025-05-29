package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/types"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	Rpc             string `yaml:"rpc"`
	StartCheckPoint string `yaml:"startCheckPoint"`
	EndCheckPoint   string `yaml:"endCheckPoint"`
	ClientPath      string `yaml:"clientPath"`
	Address         string `yaml:"address"`
	Model           int    `yaml:"model"`
}

func readConfigFile() (*Config, error) {
	_, err := os.Getwd()
	if err != nil {
		fmt.Println("can not get working path .. ", err)
		return nil, err
	}
	data, err := ioutil.ReadFile("./bfc_replay/config.yaml")
	if err != nil {
		fmt.Println("can not read file ：%v", err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("can not decode file ：%v", err)
	}
	return &config, nil
}

func QueryTransaction(rpcUrl string, address string, pre *bfc_types.TransactionDigest) (
	[]bfc_types.TransactionDigest,
	bool,
	bfc_types.TransactionDigest,
) {
	newAddress, _ := bfc_types.NewAddressFromHex(address)
	var cli, _ = client.Dial(rpcUrl)
	limit := uint(50)
	data, err := cli.QueryTransactionBlocks(
		context.TODO(),
		types.BfcTransactionBlockResponseQuery{
			Filter: &types.TransactionFilter{
				FromAddress: newAddress,
			},
		},
		pre,
		&limit,
		false,
	)
	var txList []bfc_types.TransactionDigest
	var lastDigest bfc_types.TransactionDigest
	if err != nil || data == nil || len(data.Data) == 0 {
		fmt.Println("err = or data is nil", err)
		return txList, false, lastDigest
	}
	for i := range data.Data {
		r := &data.Data[i]
		txList = append(txList, r.Digest)
		lastDigest = r.Digest
	}
	return txList, data.HasNextPage, lastDigest
}

func main() {
	config, err := readConfigFile()
	if err != nil {
		return
	}
	fmt.Printf("begin.\n")
	exec.Command(config.ClientPath, "new-env", "--rpc", config.Rpc, "--alias", "alpha")
	exec.Command(config.ClientPath, "switch", "--env", "alpha")
	if config.Model == 0 {
		cmd := exec.Command(
			config.ClientPath, "client", "replay-checkpoint",
			"--start", config.StartCheckPoint, "--end", config.EndCheckPoint,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		fmt.Printf("end.\n")
	} else {
		txList, flag, lastDigest := QueryTransaction(config.Rpc, config.Address, nil)
		for {
			for i := range txList {
				digest := txList[i]
				cmd := exec.Command(config.ClientPath, "client", "replay-transaction", "--tx-digest", digest.String())
				var out bytes.Buffer
				var errOut bytes.Buffer
				cmd.Stdout = &out
				cmd.Stderr = &errOut
				err := cmd.Run()
				if err != nil {
					fmt.Printf("Error output: %q\n", errOut.String())
				}
				if !strings.Contains(out.String(), "Status: Success") {
					fmt.Printf("Command output fail: %q\n", digest)
				}
			}
			if flag {
				txList, flag, lastDigest = QueryTransaction(config.Rpc, config.Address, &lastDigest)
			} else {
				fmt.Printf("end.\n")
				return
			}
		}
	}
}
