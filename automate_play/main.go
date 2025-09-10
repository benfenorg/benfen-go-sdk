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
	var swap_poolid = "BFC8fe191960d08ff76553f74241aaf4409509d214410e3b562fdccdaa68cc189127c9c"
	var abfc_address = "BFCb23b9de38dc09b4e56f3b34309d79fe3c735daf4328b83f3d0b559f6cbbc10fc98fe"
	var to_address = "BFC64d767329da16653c62eb6b0e85bd5b7f0fd2f325ff2bedb7f00d54d2a2de38e5133"

	var swap_poolid_usd = "BFC9285c8e5ca84049cab1f69528abe11e6e9a80821c4691386e8fb5630e2831bbfa129"
	var abfc_usd_address = "BFC92c6e273a48a60ac6ae1bd5b04678e050b761853d1fab89eb08fcf8ef67ed922bdfd"
	var packageName = "0x82422187d785f431d809afc12b4dcd6c025a6f2ad1dbcd73afe0b8a8ed903d58::anonymous_usd::ANONYMOUS_USD"

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

		cmd.GetAnonymousValueUsd(
			bfcPath,
			abfc_usd_address,
			packageName,
		)

		cmd.SplitAndTransferUsd(bfcPath, abfc_usd_address, to_address, packageName)
		time.Sleep(5 * time.Second)
	}

}
