package main

import (
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/cmd"
	"time"
)

func main() {
	var LocalNetFaucetUrl = "http://127.0.0.1:9123/gas"
	var address = "BFCfc171f86c07b0311a347d7e71b261c684848becbececec78802f1bf8a599f729d85a"

	var bfcPath = "/data/obc/bfc"
	var swap_poolid = "BFC324d07bd342e01611ba567d70432ca921aac5900ce94f6438abc729aea0929a2f904"
	var abfc_address = "BFC62cbbed60553263d1b73dc44ad94148cc1d567a774089f965aa58c751223eb83fa05"
	var to_address = "BFC64d767329da16653c62eb6b0e85bd5b7f0fd2f325ff2bedb7f00d54d2a2de38e5133"
	for {
		var cli, _ = client.Dial("http://localhost:9000")
		var objectid, err = cli.GetCoinsRpc()
		println(objectid)
		if err != nil {
			client.FaucetFundAccount(address, LocalNetFaucetUrl)
		}

		objectid, _ = cli.GetCoinsRpc()
		cmd.SwapIn(
			bfcPath,
			objectid,
			swap_poolid,
		)

		cmd.SwapOut(
			bfcPath,
			abfc_address,
			swap_poolid,
		)

		cmd.GetAnonymousValue(
			bfcPath,
			abfc_address,
		)

		cmd.SplitAndTransfer(bfcPath, abfc_address, to_address)
		time.Sleep(5 * time.Second)
	}

}
