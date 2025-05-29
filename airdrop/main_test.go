package main

import (
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/airdrop/event"
	"github.com/benfenorg/benfen-go-sdk/airdrop/transfer"
	"github.com/benfenorg/benfen-go-sdk/airdrop/util"
	"testing"
)

func Test_Apy(t *testing.T) {
	transfer.GetAddressBalance(
		"0x7419050e564485685f306e20060472fca1b3a4453b41bdace0010624801b11ea",
		"https://testrpc.benfen.org/",
	)
}

// 1711960834951
func Test_GetTime(t *testing.T) {
	time := util.GetTimeStamp()
	fmt.Printf("time: %d\n", time)

}

func Test_GetValidatorInfoEvent(t *testing.T) {
	var nodeRpc = "http://0.0.0.0:9000"

	event.QueryValidatorInfoEvent(nodeRpc, event.ValidatorInfoEvent)

}
