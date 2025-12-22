package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/benfenorg/benfen-go-sdk/bfc_types"
)

const (
	DevNetFaucetUrl   = "https://obcfaucet.openblock.vip/gas"
	TestNetFaucetUrl  = "https://obcfaucet.openblock.vip/gas"
	LocalNetFaucetUrl = "http://127.0.0.1:9123/gas"
)

type EncodeDataParams struct {
	Value int64  `json:"value"`
	Owner string `json:"owner"`
}

func GetEncodeData(url string) {
	// 准备要发送的数据
	data := EncodeDataParams{
		Value: 30,
		Owner: "0x01",
	}
	jsonData, _ := json.Marshal(data)

	// 发送POST请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 解析JSON响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("解析响应失败:", err)
		return
	}

	fmt.Println("响应结果:", result)
}

func FaucetFundAccount(address string, faucetUrl string) (string, error) {
	_, err := bfc_types.NewAddressFromHex(address)
	if err != nil {
		return "", err
	}

	paramJson := fmt.Sprintf(`{"FixedAmountRequest":{"recipient":"%v"}}`, address)
	request, err := http.NewRequest(http.MethodPost, faucetUrl, bytes.NewBuffer([]byte(paramJson)))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 && res.StatusCode != 201 {
		return "", fmt.Errorf("post %v response code = %v", faucetUrl, res.Status)
	}
	defer res.Body.Close()

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response struct {
		TransferredGasObjects []struct {
			Amount uint64 `json:"amount"`
			Id     string `json:"id"`
			Digest string `json:"transferTxDigest"`
		} `json:"transferredGasObjects,omitempty"`
		Error string `json:"error,omitempty"`
	}
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(response.Error) != "" {
		return "", errors.New(response.Error)
	}
	if len(response.TransferredGasObjects) <= 0 {
		return "", errors.New("transaction not found")
	}

	return response.TransferredGasObjects[0].Digest, nil
}
