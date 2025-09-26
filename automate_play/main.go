package main

import (
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/cmd"
	"time"
)

func main() {
	var LocalNetFaucetUrl = "http://127.0.0.1:5003/gas"
	var address = "BFCfc171f86c07b0311a347d7e71b261c684848becbececec78802f1bf8a599f729d85a"

	var bfcPath = "/data/obc/bin/bfc"
	var swap_poolid = "0xdf2451aa88809a0bd1f7bfda71083160bc7cd0634539d37cb92d52494022054f"
	var abfc_address = "0xe2b59a41283f2efe5b5ed61aae23107504cc3ae67fd68b4ae0a6cec70a05dc58"
	var to_address = "BFC64d767329da16653c62eb6b0e85bd5b7f0fd2f325ff2bedb7f00d54d2a2de38e5133"

	var swap_poolid_usd = "BFCc1ebf7720a86b9387baf64e658fc2f8f242959e4d37b58cd6c59e7c334aa5355e383"
	var abfc_usd_address = "BFC163e6d0652e0bd90460052f01499f511c849e81738b19176abdbbcdb0d5eb4d31b23"
	var packageName = "0xb4b088e4d0d0f20d559e3f72178727cb38a21762cc758b96cf2973a5f1b31e8d::anonymous_usd::ANONYMOUS_USD"

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

		cmd.SwapOutUsd(
			bfcPath,
			abfc_usd_address,
			swap_poolid_usd,
			packageName,
		)

	//	cmd.GetAnonymousValueUsd(
	//		bfcPath,
	//		abfc_usd_address,
	//		packageName,
	//	)

		cmd.SplitAndTransferUsd(bfcPath, abfc_usd_address, to_address, packageName)
		time.Sleep(5 * time.Second)
	}

}
