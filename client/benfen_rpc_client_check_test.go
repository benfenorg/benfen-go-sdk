package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/cmd"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/move_types"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/types"
	"github.com/fardream/go-bcs/bcs"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestFeatures(t *testing.T) {
	var bfcPath = "/Users/wangruoxing/code/obc/target/debug/bfc"

	var address = cmd.ActiveAddress(bfcPath)
	println("address:", address)
	addressFromHex, _ := bfc_types.NewAddressFromHex(address)
	client := LocalnetClient(t)

	// 没有找到合适的入参，查看了sui官网 bfc相关文档以及代码
	// GetDynamicFieldObject(t, client)

	GetOwnedObjects(t, client, address)
	QueryEvents(t, client)
	QueryTransactionBlocks(t, client, addressFromHex)
	ResolveNameServiceAddress(t, client)
	ResolveNameServiceNames(t, client)
	ttt, _ := bfc_types.NewAddressFromHex("BFC27d4676db71ba2e9b4fa04d3399a97baad9a1f4e7390a085e2f0c189faa21f6c4688")
	GetStakes(t, client, ttt)

	// 这里操作比较麻烦，需要  make-validator-info  become-candidate stake  join-committee之后等到下一个epoch生效之后，才能查询
	// GetStakesByIds(t, client)

	GetMoveFunctionArgTypes(t, client)
	GetNormalizedMoveFunction(t, client)
	GetNormalizedMoveModule(t, client)
	GetNormalizedMoveModulesByPackage(t, client)
	GetNormalizedMoveStruct(t, client)

	coins, _ := client.GetCoins(context.Background(), *addressFromHex, nil, nil, 1)
	digest := cmd.Transfer(bfcPath, coins.Data[0].CoinObjectId.String(), address, "1000000000")

	GetEvents(t, client, digest)
	GetObject(t, client, addressFromHex)
	MultiGetObjects(t, client, addressFromHex)
	TryGetPastObject(t, client, addressFromHex)
	TryMultiGetPastObjects(t, client, addressFromHex)
	BatchTransaction(t, client, addressFromHex)
	MergeCoins(t, client, addressFromHex)
	MoveCall(t, client, addressFromHex)
	Pay(t, client, addressFromHex)
	PayAllBFC(t, client, addressFromHex)
	PayBFC(t, client, addressFromHex)
	Publish(t, client, addressFromHex)
	RequestAddStake(t, client, addressFromHex)
	RequestWithdrawStake(t, client, addressFromHex)
	SplitCoin(t, client, addressFromHex)
	SplitCoinEqual(t, client, addressFromHex)
	TransferObject(t, client, addressFromHex)
	TransferBFC(t, client, addressFromHex)
	DevInspectTransactionBlock(t, client, addressFromHex)
	DryRunTransaction(t, client, addressFromHex)
	txbyte := cmd.TransferUnSign(bfcPath, coins.Data[0].CoinObjectId.String(), address, "1000000000")
	signature := cmd.TransferSignJson(bfcPath, coins.Data[0].CoinObjectId.String(), address, "1000000000")
	ExecuteTransactionBlock(t, client, txbyte, signature)
	GetDynamicFields(t, client, addressFromHex)
	GetAllBalances(t, client, addressFromHex)
	GetAllCoins(t, client, addressFromHex)
	GetBalance(t, client, addressFromHex)
	GetCoinMetadata(t, client)
	GetCoinsRpc(t, client)
	GetTotalSupply(t, client)
	GetCommiteeInfo(t, client)
	GetLatestSuiSystemState(t, client)
	GetReferenceGasPrice(t, client)
	GetValidatorsApy(t, client)
	GetChainIdentifier(t, client)
	GetCheckpoint(t, client)
	GetCheckpoints(t, client)
	GetLatestCheckpointSequenceNumber(t, client)
	GetProtocolConfig(t, client)
	GetTotalTransactionBlocks(t, client)
	GetTransactionBlock(t, client, digest)
	MultiGetTransactionBlocks(t, client, digest)
}

func GetDynamicFieldObject(t *testing.T, chain *Client) {
	address, err := bfc_types.NewAddressFromHex(
		"0x6337de842b9daae517401654860d7072d6fb1d78f27dd10e562acfc7e1f09b57",
	)
	name := bfc_types.DynamicFieldName{
		Type:  "0x1::ascii::String",
		Value: "00000000000000000000000000000000000000000000000000000000000000c8::busd::BUSD",
	}
	resp, err := chain.GetDynamicFieldObject(
		context.Background(),
		*address,
		name,
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetStakesByIds(t *testing.T, cli *Client) {
	owner, err := bfc_types.NewAddressFromHex("")
	stakes, err := cli.GetStakes(context.Background(), *owner)
	require.Nil(t, err)
	require.GreaterOrEqual(t, len(stakes), 1)
	stake1 := stakes[0].Stakes[0].Data
	stakeId := stake1.StakedBfcId
	stakesFromId, err := cli.GetStakesByIds(context.Background(), []bfcObjectID{stakeId})
	require.Nil(t, err)
	require.GreaterOrEqual(t, len(stakesFromId), 1)
	queriedStake := stakesFromId[0].Stakes[0].Data
	require.Equal(t, stake1, queriedStake)
	t.Log(stakesFromId)
}

func GetDynamicFields(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	parentObjectId := addressFromHex
	limit := uint(5)
	type args struct {
		ctx            context.Context
		parentObjectId bfcObjectID
		cursor         *bfcObjectID
		limit          *uint
	}
	tests := []struct {
		name    string
		args    args
		want    *types.DynamicFieldPage
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				ctx:            context.TODO(),
				parentObjectId: *parentObjectId,
				cursor:         nil,
				limit:          &limit,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := chain.GetDynamicFields(tt.args.ctx, tt.args.parentObjectId, tt.args.cursor, tt.args.limit)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetDynamicFields() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Log(got)
			},
		)
	}
}

func MultiGetTransactionBlocks(t *testing.T, chain *Client, digest string) {
	d1, err := bfc_types.NewDigest(digest)
	resp, err := chain.multiGetTransactionBlocks(
		context.Background(),
		[]bfcDigest{*d1},
		&types.BfcTransactionBlockResponseOptions{
			ShowInput:          true,
			ShowEffects:        true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
			ShowEvents:         true,
		},
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetTransactionBlock(t *testing.T, cli *Client, digest string) {
	d, err := bfc_types.NewDigest(digest)
	require.Nil(t, err)
	resp, err := cli.GetTransactionBlock(
		context.Background(), *d, types.BfcTransactionBlockResponseOptions{
			ShowInput:          true,
			ShowEffects:        true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
			ShowEvents:         true,
		},
	)
	t.Logf("the response is %#v", resp)
}

func DevInspectTransactionBlock(t *testing.T, chain *Client, senderAddress *BfcAddress) {
	packageId, module, function, err := chain.GetFunctions("0xc8::bfc_system::vault_info")
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	typeTags := []move_types.TypeTag{}
	typeArgs := []string{"0xc8::busd::BUSD"}
	for _, arg := range typeArgs {
		typeAddr, typeModule, typeFunction, _ := chain.GetFunctions(arg)
		typeTags = append(
			typeTags, move_types.TypeTag{
				Struct: &move_types.StructTag{
					Address: *typeAddr,
					Module:  typeModule,
					Name:    typeFunction,
				},
			},
		)
	}
	args := []*DevInspectArgs{
		{
			Type:    "object",
			Value:   "0xc9",
			Version: 1,
		},
	}
	arguments := []bfc_types.CallArg{}
	for _, arg := range args {
		var a bfc_types.CallArg
		if arg.Type == "object" {
			a, _ = ptb.SharedObjCallArg(arg.Value.(string), arg.Version)
		} else if arg.Type == "pure" {
			a, _ = ptb.PureCallArg(arg.Value)
		} else {
			continue
		}
		arguments = append(arguments, a)
	}
	ptb.MoveCall(*packageId, module, function, typeTags, arguments)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(*senderAddress, nil, pt, 0, 0)
	jsonBytes, err := bcs.Marshal(tx)
	txBytes := jsonBytes[1 : len(jsonBytes)-82]
	_, err = chain.DevInspectTransactionBlock(context.Background(), *senderAddress, txBytes, nil, nil)
	require.NoError(t, err)
}

func TryMultiGetPastObjects(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	coinList, err := chain.GetCoins(context.Background(), *addressFromHex, nil, nil, 1)
	objId, err := bfc_types.NewAddressFromHex(coinList.Data[0].CoinObjectId.String())
	r1 := types.GetPastObjectRequest{
		ObjectId: *objId,
		Version:  "1",
	}
	resp, err := chain.tryMultiGetPastObjects(
		context.Background(),
		[]types.GetPastObjectRequest{r1},
		nil,
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func TryGetPastObject(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	coinList, err := cli.GetCoins(context.Background(), *addressFromHex, nil, nil, 1)
	objId, err := bfc_types.NewAddressFromHex(coinList.Data[0].CoinObjectId.String())
	require.Nil(t, err)
	data, err := cli.TryGetPastObject(context.Background(), *objId, 1, nil)
	require.Nil(t, err)
	t.Log(data)
}

func GetEvents(t *testing.T, cli *Client, digest string) {
	d, err := bfc_types.NewDigest(digest)
	require.Nil(t, err)
	res, err := cli.GetEvents(context.Background(), *d)
	require.NoError(t, err)
	t.Log(res)
}

func GetTotalTransactionBlocks(t *testing.T, cli *Client) {
	res, err := cli.GetTotalTransactionBlocks(context.Background())
	require.Nil(t, err)
	t.Log(res)
}

func GetProtocolConfig(t *testing.T, chain *Client) {
	resp, err := chain.GetProtocolConfig(
		context.Background(),
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetLatestCheckpointSequenceNumber(t *testing.T, cli *Client) {
	res, err := cli.GetLatestCheckpointSequenceNumber(context.Background())
	require.Nil(t, err)
	t.Log(res)
}

func GetCheckpoints(t *testing.T, chain *Client) {
	resp, err := chain.GetCheckPoints(
		context.Background(),
		nil,
		10,
		false,
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetCheckpoint(t *testing.T, chain *Client) {
	resp, err := chain.GetCheckPoint(
		context.Background(),
		"1",
	)
	require.NoError(t, err)
	PrintJson(resp)
	var obj types.CheckPointObject
	err = json.Unmarshal([]byte(resp), &obj)
	if err != nil {
		fmt.Println("decode JSON fail:", err)
		return
	}
	fmt.Println("decode result:")
	fmt.Println("Epoch:", obj.Epoch)
	fmt.Println("SequenceNumber:", obj.SequenceNumber)
	fmt.Println("Digest:", obj.Digest)
}

func GetChainIdentifier(t *testing.T, chain *Client) {
	resp, err := chain.GetChainIdentifier(
		context.Background(),
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetValidatorsApy(t *testing.T, chain *Client) {
	resp, err := chain.GetValidatorsApy(
		context.Background(),
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetReferenceGasPrice(t *testing.T, chain *Client) {
	resp, err := chain.GetReferenceGasPrice(
		context.Background(),
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetLatestSuiSystemState(t *testing.T, cli *Client) {
	state, err := cli.GetLatestSuiSystemState(context.Background())
	require.Nil(t, err)
	t.Logf("system state = %v", state)
}

func GetCommiteeInfo(t *testing.T, chain *Client) {
	resp, err := chain.GetCommiteeInfo(
		context.Background(),
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetTotalSupply(t *testing.T, chain *Client) {
	resp, err := chain.GetTotalSupply(
		context.Background(),
		types.BFC_COIN_TYPE,
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetCoinsRpc(t *testing.T, chain *Client) {
	defaultCoinType := types.BFCoinType
	coins, err := chain.GetCoins(context.TODO(), *Address, &defaultCoinType, nil, 1)
	require.NoError(t, err)
	t.Logf("%#v", coins)
}

func GetCoinMetadata(t *testing.T, chain *Client) {
	metadata, err := chain.GetCoinMetadata(context.TODO(), types.BFCoinType)
	require.Nil(t, err)
	t.Logf("%#v", metadata)
}

func GetBalance(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	balance, err := chain.GetBalance(context.TODO(), *addressFromHex, "")
	require.NoError(t, err)
	t.Logf(
		"Coin Name: %v, Count: %v, Total: %v, Locked: %v",
		balance.CoinType, balance.CoinObjectCount,
		balance.TotalBalance.String(), balance.LockedBalance,
	)
}

func GetAllCoins(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	type args struct {
		ctx     context.Context
		address BfcAddress
		cursor  *bfcObjectID
		limit   uint
	}
	tests := []struct {
		name    string
		chain   *Client
		args    args
		want    *types.CoinPage
		wantErr bool
	}{
		{
			name:  "test case 1",
			chain: chain,
			args: args{
				ctx:     context.TODO(),
				address: *addressFromHex,
				cursor:  nil,
				limit:   3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := chain.GetAllCoins(tt.args.ctx, tt.args.address, tt.args.cursor, tt.args.limit)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetAllCoins() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("%#v", got)
			},
		)
	}
}

func GetAllBalances(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	balances, err := chain.GetAllBalances(context.TODO(), *addressFromHex)
	require.NoError(t, err)
	for _, balance := range balances {
		t.Logf(
			"Coin Name: %v, Count: %v, Total: %v, Locked: %v",
			balance.CoinType, balance.CoinObjectCount,
			balance.TotalBalance.String(), balance.LockedBalance,
		)
	}
}

func ExecuteTransactionBlock(t *testing.T, chain *Client, txbytes string, signature string) {
	signatures := []string{signature}
	resp, err := chain.executeTransactionBlock(
		context.Background(),
		txbytes,
		signatures,
		&types.BfcTransactionBlockResponseOptions{
			ShowInput:          true,
			ShowEffects:        true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
			ShowEvents:         true,
		},
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func DryRunTransaction(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)
	amount := BFC(0.01).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 0, 0)
	require.NoError(t, err)
	tx, err := cli.PayAllBFC(
		context.Background(), *signer, *signer,
		pickedCoins.CoinIds(),
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.NoError(t, err)
	resp, err := cli.DryRunTransaction(context.Background(), tx.TxBytes)
	require.Nil(t, err)
	t.Log("dry run status:", resp.Effects.Data.IsSuccess())
	t.Log("dry run error:", resp.Effects.Data.V1.Status.Error)
}

func TransferBFC(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	sender := addressFromHex
	recipient := sender
	amount := BFC(0.1).Uint64()
	gasBudget := BFC(0.1).Uint64()
	coin := GetCoins(t, cli, *sender, 1)[0]
	gasPrice := uint64(100)
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	err := ptb.TransferBFC(*recipient, &amount)
	require.NoError(t, err)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		*sender, []*bfc_types.ObjectRef{
			coin.Reference(),
		},
		pt, gasBudget, gasPrice,
	)
	txBytesBCS, err := bcs.Marshal(tx)
	require.NoError(t, err)
	txn, err := cli.TransferBFC(
		context.Background(), *sender, *recipient, coin.CoinObjectId,
		types.NewSafeBfcBigInt(amount),
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.NoError(t, err)
	txBytesRemote := txn.TxBytes.Data()
	require.Equal(t, txBytesBCS, txBytesRemote)
}

func TransferObject(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	recipient := signer
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(coins.Data), 2)
	coin := coins.Data[0]
	txn, err := cli.TransferObject(
		context.Background(), *signer, *recipient,
		coin.CoinObjectId, nil, types.NewSafeBfcBigInt(BFC(0.01).Uint64()),
	)
	require.Nil(t, err)
	simulateCheck(t, cli, txn.TxBytes, true)
}

func SplitCoinEqual(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)
	amount := BFC(0.01).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), 0, 1, 0)
	require.NoError(t, err)
	txn, err := cli.SplitCoinEqual(
		context.Background(), *signer,
		pickedCoins.Coins[0].CoinObjectId,
		types.NewSafeBfcBigInt(uint64(2)),
		nil, types.NewSafeBfcBigInt(gasBudget),
	)
	require.Nil(t, err)
	simulateCheck(t, cli, txn.TxBytes, true)
}

func SplitCoin(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)
	amount := BFC(0.01).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), 0, 1, 0)
	require.NoError(t, err)
	splitCoins := []types.SafeBfcBigInt[uint64]{types.NewSafeBfcBigInt(amount / 2)}
	txn, err := cli.SplitCoin(
		context.Background(), *signer,
		pickedCoins.Coins[0].CoinObjectId,
		splitCoins,
		nil, types.NewSafeBfcBigInt(gasBudget),
	)
	require.Nil(t, err)
	simulateCheck(t, cli, txn.TxBytes, false)
}

func RequestWithdrawStake(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	coinList, err := chain.GetCoins(context.Background(), *signer, nil, nil, 1)
	coin1, err := bfc_types.NewAddressFromHex(coinList.Data[0].CoinObjectId.String())
	resp, err := chain.RequestWithdrawStake(
		context.Background(),
		*signer,
		*coin1,
		nil,
		decimal.NewFromInt(100000000),
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func RequestAddStake(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	validator := addressFromHex
	coinList, err := chain.GetCoins(context.Background(), *signer, nil, nil, 1)
	coin1, err := bfc_types.NewAddressFromHex(coinList.Data[0].CoinObjectId.String())
	coins := []bfc_types.ObjectID{*coin1}
	resp, err := chain.RequestAddStake(
		context.Background(),
		*signer,
		coins,
		decimal.NewFromInt(10000000000),
		*validator,
		nil,
		decimal.NewFromInt(100000000),
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func Publish(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	sender := addressFromHex
	modules := []string{
		"oRzrCwYAAAADAQACBwIZCBsgAAAYYmZjX2Rhb192b3RpbmdfcG9vbF90ZXN0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		"oRzrCwYAAAAEAQACBwINCA8gBi8SAAAMYmZjX2Rhb190ZXN0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAQICAQMCAQQCAQUCAQYCAQcA",
		"oRzrCwYAAAALAQAMAgwUAyAqBEoEBU4lB3OXAQiKAmAG6gIPCvkCDgyHA0sN0gMCAAUBEAIIAg4CEQISAAEIAAAABwABAgcAAwQEAAUDAgAABgABAAA" +
			"JAAEAAAsCAQABEwMEAAIHBgEBAwMNAAkABBELAQEIBQ8IBwAEBQYKAQcIBAABBwgAAQoCAQgCAQgBAQkAAQUBBggEAQgDAQgAAgkABQpDb3VudEV2ZW50B" +
			"0NvdW50ZXIGU3RyaW5nCVR4Q29udGV4dANVSUQHY291bnRlcgtjcmVhdGVFdmVudARlbWl0BWV2ZW50CmdldENvdW50ZXICaWQEaW5jcgRuYW1lA25ldwZv" +
			"YmplY3QGc2VuZGVyBnN0cmluZwh0cmFuc2Zlcgp0eF9jb250ZXh0BHV0ZjgFdmFsdWUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIKAgwLaGVsbG8gd29ybGQAAgIKCAMUAwECAQwIAgABBAABBQc" +
			"AEQMSATgAAgEBBAAHCwoALhEHDAELABEFBgAAAAAAAAAAEgALATgBAgIBBAABCQoAEAAUBgEAAAAAAAAAFgsADwAVAgABAA==",
	}
	dep1, err := bfc_types.NewAddressFromHex("0x0000000000000000000000000000000000000000000000000000000000000001")
	dep2, err := bfc_types.NewAddressFromHex("0x0000000000000000000000000000000000000000000000000000000000000002")
	dependencies := []bfc_types.ObjectID{*dep1, *dep2}
	r, err := chain.publish(
		context.Background(),
		sender,
		modules,
		dependencies,
		nil,
		"1000000000",
	)
	println(r)
	require.NoError(t, err)
}

func PayBFC(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	recipient := addressFromHex
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)
	amount := BFC(0.001).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 0, 0)
	require.NoError(t, err)
	txn, err := cli.PayBFC(
		context.Background(), *signer,
		pickedCoins.CoinIds(),
		[]BfcAddress{*recipient},
		[]types.SafeBfcBigInt[uint64]{
			types.NewSafeBfcBigInt(amount),
		},
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.Nil(t, err)
	simulateCheck(t, cli, txn.TxBytes, true)
}

func PayAllBFC(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	sender := addressFromHex
	recipient := sender
	gasBudget := BFC(0.1).Uint64()
	coins := GetCoins(t, cli, *sender, 2)
	coin, coin2 := coins[0], coins[1]
	gasPrice := uint64(100)
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	err := ptb.PayAllBFC(*recipient)
	require.NoError(t, err)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		*sender, []*bfc_types.ObjectRef{
			coin.Reference(),
			coin2.Reference(),
		},
		pt, gasBudget, gasPrice,
	)
	txBytesBCS, err := bcs.Marshal(tx)
	require.NoError(t, err)
	txn, err := cli.PayAllBFC(
		context.Background(), *sender, *recipient,
		[]bfcObjectID{
			coin.CoinObjectId, coin2.CoinObjectId,
		},
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.NoError(t, err)
	txBytesRemote := txn.TxBytes.Data()
	require.Equal(t, txBytesBCS, txBytesRemote)
}

func Pay(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	sender := addressFromHex
	recipient2, _ := bfc_types.NewAddressFromHex("0x123456")
	amount := BFC(0.1).Uint64()
	gasBudget := BFC(0.1).Uint64()
	coins := GetCoins(t, cli, *sender, 2)
	coin, gas := coins[0], coins[1]
	gasPrice := uint64(100)
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	err := ptb.Pay(
		[]*bfc_types.ObjectRef{coin.Reference()},
		[]BfcAddress{*recipient2, *recipient2},
		[]uint64{amount, amount},
	)
	require.NoError(t, err)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		*sender, []*bfc_types.ObjectRef{
			gas.Reference(),
		},
		pt, gasBudget, gasPrice,
	)
	txBytesBCS, err := bcs.Marshal(tx)
	require.NoError(t, err)
	resp := simulateCheck(t, cli, txBytesBCS, true)
	gasfee := resp.Effects.Data.GasFee()
	t.Log(gasfee)
}

func MoveCall(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	packageId, _ := bfc_types.NewAddressFromHex("0xc8")
	_, err := chain.MoveCall(
		context.Background(),
		*signer,
		*packageId,
		"bfc_system",
		"vault_info",
		[]string{"0xc8::busd::BUSD"},
		[]any{"0xc9"},
		nil,
		types.NewSafeBfcBigInt(uint64(100000000)),
	)
	require.NoError(t, err)
}

func MergeCoins(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)
	require.True(t, len(coins.Data) >= 3)
	coin1 := coins.Data[0]
	coin2 := coins.Data[1]
	coin3 := coins.Data[2] // gas coin
	gasBudge := BFC(1).Uint64()
	txn, err := cli.MergeCoins(
		context.Background(), *signer,
		coin1.CoinObjectId,
		coin2.CoinObjectId,
		&coin3.CoinObjectId,
		types.NewSafeBfcBigInt(gasBudge),
	)
	require.Nil(t, err)
	simulateCheck(t, cli, txn.TxBytes, true)
}

func BatchTransaction(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	signer := addressFromHex
	coins, err := chain.GetCoins(context.TODO(), *addressFromHex, nil, nil, 100)
	coin1, err := bfc_types.NewAddressFromHex(coins.Data[0].CoinObjectId.String())
	data := types.TransferObjectParams{
		ObjectId:  *coin1,
		Recipient: *signer,
	}
	transferParamMap := map[string]interface{}{
		"transferObjectRequestParams": data,
	}
	transferMaps := []map[string]interface{}{transferParamMap}
	_, err = chain.BatchTransaction(
		context.Background(),
		*signer,
		transferMaps,
		nil,
		"100000000",
	)
	require.NoError(t, err)
}

func MultiGetObjects(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	coins, err := chain.GetCoins(context.TODO(), *addressFromHex, nil, nil, 1)
	require.NoError(t, err)
	if len(coins.Data) == 0 {
		t.Log("Warning: No Object Id for test.")
		return
	}
	obj := coins.Data[0].CoinObjectId
	objs := []bfcObjectID{obj, obj}
	resp, err := chain.MultiGetObjects(
		context.Background(), objs, &types.BfcObjectDataOptions{
			ShowType:                true,
			ShowOwner:               true,
			ShowContent:             true,
			ShowDisplay:             true,
			ShowBcs:                 true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	)
	require.Nil(t, err)
	require.Equal(t, len(objs), len(resp))
	require.Equal(t, resp[0], resp[1])
}

func GetObject(t *testing.T, chain *Client, addressFromHex *BfcAddress) {
	type args struct {
		ctx   context.Context
		objID bfcObjectID
	}
	coins, err := chain.GetCoins(context.TODO(), *addressFromHex, nil, nil, 100)
	require.NoError(t, err)
	tests := []struct {
		name    string
		chain   *Client
		args    args
		want    int
		wantErr bool
	}{
		{
			name:  "test for devnet",
			chain: chain,
			args: args{
				ctx:   context.TODO(),
				objID: coins.Data[0].CoinObjectId,
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				_, err := tt.chain.GetObject(
					tt.args.ctx, tt.args.objID, &types.BfcObjectDataOptions{
						ShowType:                true,
						ShowOwner:               true,
						ShowContent:             true,
						ShowDisplay:             true,
						ShowBcs:                 true,
						ShowPreviousTransaction: true,
						ShowStorageRebate:       true,
					},
				)
				require.NoError(t, err)
			},
		)
	}
}

func GetNormalizedMoveStruct(t *testing.T, cli *Client) {
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")
	_, err = cli.getNormalizedMoveStruct(
		context.Background(),
		objId,
		"bfc_system",
		"BfcSystemState",
	)
	require.NoError(t, err)
}

func GetNormalizedMoveModulesByPackage(t *testing.T, cli *Client) {
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")
	_, err = cli.getNormalizedMoveModulesByPackage(
		context.Background(),
		objId,
	)
	require.NoError(t, err)
}

func GetNormalizedMoveModule(t *testing.T, cli *Client) {
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")
	_, err = cli.getNormalizedMoveModule(
		context.Background(),
		objId,
		"bfc_system",
	)
	require.NoError(t, err)
}

func GetNormalizedMoveFunction(t *testing.T, cli *Client) {
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")
	resp, err := cli.getNormalizedMoveFunction(
		context.Background(),
		objId,
		"bfc_system",
		"create_stake_manager_key",
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetMoveFunctionArgTypes(t *testing.T, cli *Client) {
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")
	resp, err := cli.getMoveFunctionArgTypes(
		context.Background(),
		objId,
		"bfc_system",
		"create_stake_manager_key",
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func GetStakes(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	_, err := cli.GetStakes(context.Background(), *addressFromHex)
	require.Nil(t, err)
}

func ResolveNameServiceNames(t *testing.T, cli *Client) {
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")
	resp, err := cli.resolveNameServiceNames(
		context.Background(),
		objId,
		nil,
		10,
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func ResolveNameServiceAddress(t *testing.T, cli *Client) {
	resp, err := cli.resolveNameServiceAddress(
		context.Background(),
		"example.bfc",
	)
	require.NoError(t, err)
	PrintJson(resp)
}

func QueryTransactionBlocks(t *testing.T, cli *Client, addressFromHex *BfcAddress) {
	limit := uint(10)
	type args struct {
		ctx             context.Context
		query           types.BfcTransactionBlockResponseQuery
		cursor          *bfcDigest
		limit           *uint
		descendingOrder bool
	}
	tests := []struct {
		name    string
		args    args
		want    *types.TransactionBlocksPage
		wantErr bool
	}{
		{
			name: "test for queryTransactionBlocks",
			args: args{
				ctx: context.TODO(),
				query: types.BfcTransactionBlockResponseQuery{
					Filter: &types.TransactionFilter{
						FromAddress: addressFromHex,
					},
					Options: &types.BfcTransactionBlockResponseOptions{
						ShowInput:   true,
						ShowEffects: true,
					},
				},
				cursor:          nil,
				limit:           &limit,
				descendingOrder: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := cli.QueryTransactionBlocks(
					tt.args.ctx,
					tt.args.query,
					tt.args.cursor,
					tt.args.limit,
					tt.args.descendingOrder,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("QueryTransactionBlocks() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("%#v", got)
			},
		)
	}
}

func QueryEvents(t *testing.T, cli *Client) {
	limit := uint(10)
	type args struct {
		ctx             context.Context
		query           types.EventFilter
		cursor          *types.EventId
		limit           *uint
		descendingOrder bool
	}
	tests := []struct {
		name    string
		args    args
		want    *types.EventPage
		wantErr bool
	}{
		{
			name: "test for query events",
			args: args{
				ctx: context.TODO(),
				query: types.EventFilter{
					Sender: CrossChainAddress,
				},
				cursor:          nil,
				limit:           &limit,
				descendingOrder: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := cli.QueryEvents(
					tt.args.ctx,
					tt.args.query,
					tt.args.cursor,
					tt.args.limit,
					tt.args.descendingOrder,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("QueryEvents() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Log(got)
			},
		)
	}
}

func GetOwnedObjects(t *testing.T, cli *Client, address string) {
	obj, err := bfc_types.NewAddressFromHex("0x2")
	addressFromHex, err := bfc_types.NewAddressFromHex(address)
	require.Nil(t, err)
	query := types.BfcObjectResponseQuery{
		Filter: &types.BfcObjectDataFilter{
			Package: obj,
		},
		Options: &types.BfcObjectDataOptions{
			ShowType: true,
		},
	}
	limit := uint(1)
	objs, err := cli.GetOwnedObjects(context.Background(), *addressFromHex, &query, nil, &limit)
	require.Nil(t, err)
	require.GreaterOrEqual(t, len(objs.Data), int(limit))
}
