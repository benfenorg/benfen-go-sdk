package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/types"
)

func convertBfcAddressToHexAddress(address string) string {

	length := len(address)
	address = address[3 : length-4]

	address = "0x" + address
	return address
}

type MyObject struct {
	Result []json.RawMessage `json:"result"`
}

type NameType struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PoolObject struct {
	Name       NameType `json:"name"`
	BcsName    string   `json:"bcsName"`
	Type       string   `json:"type"`
	ObjectType string   `json:"objectType"`
	ObjectID   string   `json:"objectId"`
	Version    int      `json:"version"`
	Digest     string   `json:"digest"`
}

func ConvertPoolObject(resp string) ([]PoolObject, error) {
	var result [][]json.RawMessage
	var data []PoolObject

	err := json.Unmarshal([]byte(resp), &result)
	if err != nil {
		fmt.Println("Error 1 :", err)
		return nil, err
	}

	data = make([]PoolObject, len(result))
	for index := range result {
		var innerObj PoolObject
		err := json.Unmarshal(result[index][1], &innerObj)
		if err != nil {
			fmt.Println("Error 2:", err)
			return nil, err
		}

		fmt.Printf("Name: %s\n", innerObj.Name.Value)
		fmt.Printf("BcsName: %s\n", innerObj.BcsName)
		fmt.Printf("Type: %s\n", innerObj.Type)
		fmt.Printf("ObjectType: %s\n", innerObj.ObjectType)
		fmt.Printf("ObjectID: %s\n", innerObj.ObjectID)
		fmt.Printf("Version: %d\n", innerObj.Version)
		fmt.Printf("Digest: %s\n", innerObj.Digest)

		data[index] = innerObj
	}

	return data, nil
}

func ConvertToNftDetailInfo(resp string) (NftDetailInfo, error) {
	var data NftDetailInfo
	err := json.Unmarshal([]byte(resp), &data)
	if err != nil {
		fmt.Println("Error 1 :", err)
		return NftDetailInfo{}, err
	}

	return data, nil

}

func ConvertToStablePoolDetailInfo(resp string) (StablePoolDetailInfo, error) {
	var data StablePoolDetailInfo
	err := json.Unmarshal([]byte(resp), &data)
	if err != nil {
		fmt.Println("Error 1 :", err)
		return StablePoolDetailInfo{}, err
	}

	return data, nil

}

type NftDetailInfo struct {
	Data NftDetailInfoData `json:"data"`
}

type NftDetailInfoData struct {
	Content NftDetailInfoContent `json:"content"`
}

type NftDetailInfoContent struct {
	Fields NftDetailInfoFields `json:"fields"`
}

type NftDetailInfoFields struct {
	Name  string                  `json:"name"`
	Value NftDetailInfoDataValues `json:"value"`
}

type NftDetailInfoDataValues struct {
	Fields NftDetailInfoDataValuesFields `json:"fields"`
}

type NftDetailInfoDataValuesFields struct {
	BfcBalance string `json:"bfc_balance"`
	TotalPower string `json:"total_power"`
	Period     string `json:"period"`
}

type StablePoolDetailInfo struct {
	Data StablePoolDetailInfoData `json:"data"`
}

type StablePoolDetailInfoData struct {
	Content StablePoolDetailInfoContent `json:"content"`
}
type StablePoolDetailInfoContent struct {
	Fields StablePoolDetailInfoFields `json:"fields"`
}
type StablePoolDetailInfoFields struct {
	Name  string                         `json:"name"`
	Value StablePoolDetailInfoDataValues `json:"value"`
}
type StablePoolDetailInfoDataValues struct {
	Fields StablePoolDetailInfoDataValuesFields `json:"fields"`
}
type StablePoolDetailInfoDataValuesFields struct {
	Stable_balance string `json:"stable_balance"`
	Rewards_pool   string `json:"rewards_pool"`
}

func GetStablePoolsInfo(
	poolIdString string,
	stablePoolsId string, chain *client.Client,
) (StablePoolDetailInfo, error) {
	hexAddress := convertBfcAddressToHexAddress(poolIdString)
	objectID, _ := bfc_types.NewAddressFromHex(hexAddress)
	got, err := chain.GetObjectString(
		context.Background(), *objectID, &types.BfcObjectDataOptions{
			ShowType:                true,
			ShowOwner:               true,
			ShowContent:             true,
			ShowDisplay:             true,
			ShowBcs:                 true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	)

	if err != nil {
		println("get stablePools detail info failed", err)
	}

	//println("the Content is ", got)

	stablePoolDetailInfo, err := ConvertToStablePoolDetailInfo(got)

	println("the stable pools id is", stablePoolsId)
	println("the pool obj id is ", poolIdString)
	println("the name is ", stablePoolDetailInfo.Data.Content.Fields.Name)
	println("the Stable_balance is ", stablePoolDetailInfo.Data.Content.Fields.Value.Fields.Stable_balance)
	println("the Rewards_pool is ", stablePoolDetailInfo.Data.Content.Fields.Value.Fields.Rewards_pool)

	return stablePoolDetailInfo, err

}
