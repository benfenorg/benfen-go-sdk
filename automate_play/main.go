package main

import (
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/cmd"
	"time"
)

func main() {
	var LocalNetFaucetUrl = "http://127.0.0.1:5003/gas"
	var address = "BFCfc171f86c07b0311a347d7e71b261c684848becbececec78802f1bf8a599f729d85a"

	var bfcPath = "/data/obc/bfc"
	var swap_poolid = "0x1a31c5422304ecf8ce00236aba8db9a8b952089d47828e59ab84206bd9cb23f8"
	var abfc_address = "BFC0e2c0ca5c88c3903174393975ed354cbbb4cd0052051ad3ddb06bec86f980b8e8e2c"
	var to_address = "BFC64d767329da16653c62eb6b0e85bd5b7f0fd2f325ff2bedb7f00d54d2a2de38e5133"

	var swap_poolid_usd = "BFCc42d4dc6b7f091da794467685e036262d382e5a557f380a77c30cfa11951a5f65e06"
	var abfc_usd_address = "BFC90f01b007266d2de7f5cb944bfc0b5cd413d832fee19648b4515b8c0b4b88a93895b"
	var packageName = "0x855c02088334ec40a7d103d145a95b824be9037a77919e6daf8dd3b2521f42c0::ausd::AUSD"

	for {
		var cli, _ = client.Dial("http://localhost:9000")
		var objectid, err = cli.GetCoinsRpc()
		//println(objectid)
		println("result:" ,objectid, "err", err)
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

		//cmd.GetAnonymousValue(
		//	bfcPath,
		//		abfc_address,
		//	)

		cmd.SplitAndTransfer(bfcPath, abfc_address, to_address)
		time.Sleep(5 * time.Second)

		objectid, err = cli.GetCoinsRpc()
		println(objectid)
		if err != nil {
			client.FaucetFundAccount(address, LocalNetFaucetUrl)
		}

		objectid, _ = cli.GetCoinsRpc()
		cmd.SwapInUsd(
			bfcPath,
			objectid,
			swap_poolid_usd,
			packageName,
		)
		//
		cmd.SwapOutUsd(
			bfcPath,
			abfc_usd_address,
			swap_poolid_usd,
		packageName,
		)
		//
		//
		//cmd.SplitAndTransferUsd(bfcPath, abfc_usd_address, to_address, packageName)
		time.Sleep(5 * time.Second)
	}

}
