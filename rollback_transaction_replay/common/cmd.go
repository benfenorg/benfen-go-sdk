package common

import (
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/client"
	"github.com/benfenorg/benfen-go-sdk/types"
	"golang.org/x/net/context"
	"os/exec"
	"strings"
)

func DecodeRawTransaction(input string, basePath string) string {

	cmd := exec.Command("ls", "-l") // 要执行的命令和参数

	output, err := cmd.Output() // 执行命令并获取输出

	if err != nil {
		fmt.Println("命令执行失败:", err)
		return ""
	}

	fmt.Println(string(output)) // 输出命令执行结果

	cmd = exec.Command(basePath, "--help") // 要执行的命令和参数
	output, err = cmd.Output()             // 执行命令并获取输出

	if err != nil {
		fmt.Println("命令执行失败:", err)
		return ""
	}

	fmt.Println(string(output)) // 输出命令执行结果

	cmd = exec.Command(
		basePath,
		"keytool",
		"decode-raw-transaction",
		"--tx-bytes",
		input,
	) // 要执行的命令和参数
	output, err = cmd.CombinedOutput() // 执行命令并获取输出
	//fmt.Println(string(output))        // 输出命令执行结果
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return ""
	}

	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	}
	return strings.TrimSpace(result[1])
}

func IsSystemContract(cli *client.Client, digest string, ctx context.Context) bool {

	d, _ := bfc_types.NewDigest(digest)
	options := types.BfcTransactionBlockResponseOptions{
		ShowInput: true,
		//ShowEffects: true,
	}
	respstr, _ := cli.GetTransactionBlockString(
		ctx,
		*d,
		options,
	)

	if strings.Contains(
		respstr,
		"\"sender\":\"BFC000000000000000000000000000000000000000000000000000000000000000060e0\"",
	) {
		return true
	}
	return false
}
