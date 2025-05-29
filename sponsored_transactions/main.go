package main

import (
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/sponsored_transactions/util"
)

func main() {
	sender := "0x7419050e564485685f306e20060472fca1b3a4453b41bdace0010624801b11ea"
	sponor := "0xfc171f86c07b0311a347d7e71b261c684848becbececec78802f1bf8a599f729"
	senderMnemonic := "xxxx"
	sponorMnemonic := "xxxx"

	util.TransferObjectAllowSponsor(
		sender,
		sponor,
		senderMnemonic,
		sponorMnemonic,
		uint64(1000),
		"https://testrpc.benfen.org/",
	)
	fmt.Printf("Hello, world!\n")
}
