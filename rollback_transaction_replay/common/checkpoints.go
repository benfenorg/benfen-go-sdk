package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/types"
	"strconv"
)

func GetAllCheckPoints(cli *client.Client, epochNum string) (int64, int64, error) {
	context := context.Background()
	//require.NoError(t, err)
	resp, err := cli.GetEpochs(
		context,
		epochNum,
		1,
	)
	//var result = &types.EpochPage{}
	//err = jsoniter.UnmarshalFromString(resp, result)
	fmt.Println(resp.Data[0].FirstCheckpointId, ", ", resp.Data[0].EndOfEpochInfo.LastCheckpointId)
	if err != nil {
		println("err = ", err.Error())
		return 0, 0, err
	}
	firstCheckpointId, _ := strconv.ParseInt(resp.Data[0].FirstCheckpointId, 10, 64)
	endCheckpointId, _ := strconv.ParseInt(resp.Data[0].EndOfEpochInfo.LastCheckpointId, 10, 64)
	return firstCheckpointId, endCheckpointId, nil
}

func GetTransactionFromCheckpoints(cli *client.Client, checkPoint string) ([]string, error) {
	context := context.Background()
	//require.NoError(t, err)
	resp, err := cli.GetCheckPoint(
		context,
		checkPoint,
	)
	if err != nil {
		println("err = ", err.Error())
		return nil, err
	}

	PrintJson(resp)

	var obj types.CheckPointObject
	err = json.Unmarshal([]byte(resp), &obj)
	if err != nil {
		fmt.Println("decode JSON fail:", err)
		return nil, err
	}

	fmt.Println("decode result:")
	fmt.Println("Epoch:", obj.Epoch)
	fmt.Println("SequenceNumber:", obj.SequenceNumber)
	fmt.Println("Digest:", obj.Digest)
	fmt.Println("NetworkTotalTransactions:", obj.NetworkTotalTransactions)
	fmt.Println("PreviousDigest:", obj.PreviousDigest)
	fmt.Println("TimestampMs:", obj.TimestampMs)
	fmt.Println("Transactions:", obj.Transactions)
	fmt.Println("CheckpointCommitments:", obj.CheckpointCommitments)
	fmt.Println("ValidatorSignature:", obj.ValidatorSignature)
	fmt.Println("EpochRollingGasCostSummary:")
	fmt.Println("  ComputationCost:", obj.EpochRollingGasCostSummary.ComputationCost)
	fmt.Println("  StorageCost:", obj.EpochRollingGasCostSummary.StorageCost)
	fmt.Println("  StorageRebate:", obj.EpochRollingGasCostSummary.StorageRebate)
	fmt.Println("  NonRefundableStorageFee:", obj.EpochRollingGasCostSummary.NonRefundableStorageFee)

	return obj.Transactions, nil
}
