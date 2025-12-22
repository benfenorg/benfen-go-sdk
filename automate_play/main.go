package main

import (
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/cmd"
	"math/rand"
	"time"
)

func main2() {
	var LocalNetFaucetUrl = "http://127.0.0.1:5003/gas"
	var address = "BFCfc171f86c07b0311a347d7e71b261c684848becbececec78802f1bf8a599f729d85a"

	var bfcPath = "/data/obc/bfc"
	var swap_poolid = "0x1a31c5422304ecf8ce00236aba8db9a8b952089d47828e59ab84206bd9cb23f8"
	var abfc_address = "BFC0e2c0ca5c88c3903174393975ed354cbbb4cd0052051ad3ddb06bec86f980b8e8e2c"
	//var to_address = "BFC64d767329da16653c62eb6b0e85bd5b7f0fd2f325ff2bedb7f00d54d2a2de38e5133"

	var swap_poolid_usd = "BFCc42d4dc6b7f091da794467685e036262d382e5a557f380a77c30cfa11951a5f65e06"
	var abfc_usd_address = "BFCaf95f48e6dc622edfb2e96ee0aaef8f3b6cee7e3f80f69832a201a439d5bed49fec4"
	var packageName = "0x855c02088334ec40a7d103d145a95b824be9037a77919e6daf8dd3b2521f42c0::ausd::AUSD"

	for {
		var cli, _ = client.Dial("http://localhost:9000")
		var objectid, err = cli.GetCoinsRpc()
		//println(objectid)
		println("result:", objectid, "err", err)
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

		//cmd.SplitAndTransfer(bfcPath, abfc_address, to_address)
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

func main() {
	var bfcPath = "/data/obc/bfc"
	to_address := [5]string{
		"BFC059f50e6c77ebb9152ec4eff701291a9bd9ae784947400a5da334f37c94d3496858a",
		"BFC834b0191a5eed474d032cc85c1faa9479c7cc045a64fb599fcec8d9d869c9483a27b",
		"BFC8497f504cd14ec4fb8cf6e68bd326991db70355e4bbce0f09ddd41f2ac009f8be9f9",
		"BFC9b4b325b7cf9fe67d0edfc4c4d8535f9a4a23bc1701c53105521978972fd812f911a",
		"BFCa4d931b5bdb5e5ac8c422e24d4ed87af37245dc6193db67396929dfb7864c58f5605",
	}

	var abfc_address = "BFCf2fde0dacc8955897da425e45a430b257c336e32d2e723b724a0ef19c3a6e3102eda"

	for i := 1; i <= 40; i++ {
		for _, address := range to_address {
			rand.Seed(time.Now().UnixNano())
			randomNum := rand.Int63n(201)
			encode1, encode2 := client.GetEncodeData("http://127.0.0.1:9010/rpc_internal", randomNum*1000000000)
			cmd.SplitAndTransfer(bfcPath, abfc_address, address, encode1, encode2)
			time.Sleep(1 * time.Second)
		}
	}
}
