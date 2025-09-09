package main

import (
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/cmd"
	"time"
)

func main() {
	var LocalNetFaucetUrl = "http://127.0.0.1:9123/gas"
	var address = "BFCfc171f86c07b0311a347d7e71b261c684848becbececec78802f1bf8a599f729d85a"

	var bfcPath = "/Users/wubin/workspace/suitest/obc/target/debug/bfc"
	var swap_poolid = "BFC1e2aedf8d38c3287b18e30d005429961e06f15cf5493decc501cab672995380f2fe1"
	var abfc_usd_address = "BFC3e0a5b55e6cd9542311b93445198ca3fce3f3c5b7eca25d3f5cf0761407dc734c4f1"
	var to_address = "BFC64d767329da16653c62eb6b0e85bd5b7f0fd2f325ff2bedb7f00d54d2a2de38e5133"
	var packageName = "0x18ab8dec48b7f57e374e5a6c60ea9f1183acad05ef9044dc733cd057c4d0518f::anonymous_usd::ANONYMOUS_USD"
	for {
		var cli, _ = client.Dial("http://localhost:9000")
		var objectid, err = cli.GetCoinsRpc()
		println(objectid)
		if err != nil {
			client.FaucetFundAccount(address, LocalNetFaucetUrl)
		}

		objectid, _ = cli.GetCoinsRpc()
		cmd.SwapInUsd(
			bfcPath,
			objectid,
			swap_poolid,
			packageName,
		)

		cmd.SwapOutUsd(
			bfcPath,
			abfc_usd_address,
			swap_poolid,
			packageName,
		)

		cmd.GetAnonymousValueUsd(
			bfcPath,
			abfc_usd_address,
			packageName,
		)

		cmd.SplitAndTransferUsd(bfcPath, abfc_usd_address, to_address, packageName)
		time.Sleep(5 * time.Second)
	}

}
