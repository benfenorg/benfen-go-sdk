package main

import (
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/account"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Node struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Nodename string `yaml:"nodename"`
		Password string `yaml:"password"`
		Mnemonic string `yaml:"mnemonic"`
	}
}

func ReadConfigFile() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("can not get working path .. ", err)
		return
	}

	fmt.Println("current working path：", wd)

	data, err := ioutil.ReadFile("./crosschain_node/config/node5.yaml")
	if err != nil {
		log.Fatalf("can not read file ：%v", err)
	}

	// parse config file.
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("can not decode file ：%v", err)
	}

	// using config file
	fmt.Println("Host: ", config.Node.Host)
	fmt.Println("Port: ", config.Node.Port)
	fmt.Println("nodename: ", config.Node.Nodename)
	fmt.Println("password: ", config.Node.Password)
	fmt.Println("mnemonic: ", config.Node.Mnemonic)

	account, err := account.NewAccountWithMnemonic(config.Node.Mnemonic)
	fmt.Println("address from Mnemonic: ", account.Address)
}
