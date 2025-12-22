package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

type JSONRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Result1   string `json:"result1"`
		Result2   string `json:"result2"`
		Operation string `json:"operation"`
		Timestamp int64  `json:"timestamp"`
	} `json:"result"`
	Error interface{} `json:"error"` // 可以是null或错误对象
}

func GetEncodeData(url string, value int64) (string, string) {
	// 构建请求体
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "bfcx_getAnonymousEncodeData", // 这里填写你的方法名
		"params": map[string]interface{}{
			"value": value,
			"owner": "0x1", // 这里填写owner值
		},
		"id": 5,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	fmt.Printf("请求体: %s\n", string(jsonData))

	// 发送请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", ""
	}
	defer resp.Body.Close()

	// 读取响应
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("响应: %s\n", string(body))

	var result JSONRPCResponse
	if err := json.Unmarshal(body, &result); err != nil {
		println("解析失败")
		return "", ""
	}

	return result.Result.Result1, result.Result.Result2
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
