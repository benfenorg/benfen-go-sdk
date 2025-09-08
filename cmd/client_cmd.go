package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func GetFromFaucet(basePath, myAddress string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"faucet",
	)
	output, err := cmd.CombinedOutput()
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("命令执行失败:", err.Error())
		return ""
	} else {
		fmt.Println("GetFromFaucet命令执行成功")
	}
	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	}
	return strings.TrimSpace(result[0])
}

func SplitAndTransfer(basePath, abfc_address, toAdress string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"call",
		"--function",
		"split_and_transfer",
		"--module",
		"anonymous_pay",
		"--package",
		"BFC000000000000000000000000000000000000000000000000000000000000000268e4",
		"--type-args",
		"0x2::abfc::ABFC",
		"--args",
		abfc_address,
		"1000",
		toAdress,
		"--gas-budget",
		"100000000",
	)
	output, err := cmd.CombinedOutput()
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("SplitAndTransfer命令执行失败:", err.Error())
		return ""
	} else {
		fmt.Println("命令执行成功")
	}
	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	}
	return strings.TrimSpace(result[0])
}

func SwapIn(basePath, coinId, swapPoolId string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"call",
		"--function",
		"swap_in",
		"--module",
		"anonymous_coin",
		"--package",
		"BFC000000000000000000000000000000000000000000000000000000000000000268e4",
		"--type-args",
		"0x2::abfc::ABFC",
		"0x2::bfc::BFC",
		"--args",
		coinId,
		swapPoolId,
		"--gas-budget",
		"100000000",
	)
	output, err := cmd.CombinedOutput()
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("SwapIn命令执行失败:", err.Error(), "coinId:", coinId, "swappoolId:", swapPoolId)

		return ""
	} else {
		fmt.Println("命令执行成功")
	}
	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	}
	return strings.TrimSpace(result[0])
}

func SwapOut(basePath, AbfccoinId, swapPoolId string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"call",
		"--function",
		"swap_out_with_amount",
		"--module",
		"anonymous_coin",
		"--package",
		"BFC000000000000000000000000000000000000000000000000000000000000000268e4",
		"--type-args",
		"0x2::abfc::ABFC",
		"0x2::bfc::BFC",
		"--args",
		AbfccoinId,
		"100000",
		swapPoolId,
		"--gas-budget",
		"100000000",
	)
	output, err := cmd.CombinedOutput()
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("SwapOut命令执行失败:", err.Error())
		return ""
	} else {
		fmt.Println("命令执行成功")
	}
	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	}
	return strings.TrimSpace(result[0])
}

func GetAnonymousValue(basePath, AbfccoinId string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"call",
		"--function",
		"get_anonymous_value",
		"--module",
		"anonymous_coin",
		"--package",
		"BFC000000000000000000000000000000000000000000000000000000000000000268e4",
		"--type-args",
		"0x2::abfc::ABFC",
		"--args",
		AbfccoinId,
		"0f31177f8ece16b2cfb8c1ba0b71f73252acaa6cfbbe13d36c3320617f05bc7f9a860f16c8b10c787455a01ca7bcca3469858aae4e369bc994ab64967f1fd20f",
		"8496d3d932986b43bb64b5d5c7548d5c97a73aebf4301447f3746680b2114ae1",
		"--gas-budget",
		"100000000",
	)
	output, err := cmd.CombinedOutput()
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("GetAnonymousValue命令执行失败:", err.Error())
		return ""
	} else {
		fmt.Println("命令执行成功")
	}
	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	}
	return strings.TrimSpace(result[0])
}

func ActiveAddress(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"active-address",
	)
	output, err := cmd.CombinedOutput()
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return ""
	}
	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	} else {
		fmt.Println("命令执行成功")
	}
	return strings.TrimSpace(result[0])
}

func ActiveEnv(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"active-env",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Address(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"addresses",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Gas(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"gas",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Object(basePath string, obj string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"object",
		obj,
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Balance(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"balance",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func ChainIdentifier(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"chain-identifier",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Envs(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"envs",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func NewAddress(basePath string, t string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"new-address",
		t,
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Objects(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"objects",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func MergeCoin(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"merge-coin",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func DynamicField(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"dynamic-field",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func SplitCoin(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"split-coin",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Pay(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"pay",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func PayBfc(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"pay-bfc",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Transfer(basePath string, objectId string, to string, budget string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"transfer",
		"--object-id",
		objectId,
		"--to",
		to,
		"--gas-budget",
		budget,
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return ""
	}
	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	}
	re := regexp.MustCompile(`Transaction Digest: (\S+)`)
	match := re.FindStringSubmatch(result[0])
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

func TransferUnSign(basePath string, objectId string, to string, budget string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"transfer",
		"--object-id",
		objectId,
		"--to",
		to,
		"--gas-budget",
		budget,
		"--serialize-unsigned-transaction",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return ""
	}
	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	}
	return strings.Replace(strings.Replace(result[0], "\r", "", -1), "\n", "", -1)
}

func TransferSignJson(basePath string, objectId string, to string, budget string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"transfer",
		"--object-id",
		objectId,
		"--to",
		to,
		"--gas-budget",
		budget,
		"--serialize-signed-transaction",
		"--json",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	result := strings.Split(string(output), "│")
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return ""
	}
	if len(result) < 1 {
		fmt.Println("命令返回参数失败")
		return ""
	}
	re := regexp.MustCompile(`"tx_signatures":\s*\[\s*"([^"]+)"\s*\]`)
	matches := re.FindStringSubmatch(result[0])
	return matches[1]
}

func TransferBfc(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"transfer-bfc",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Switch(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"switch",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func ProfileTransaction(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"profile-transaction",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func TxBlock(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"tx-block",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Call(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"call",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func VerifyBytecodeMeter(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"verify-bytecode-meter",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func VerifySource(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"verify-source",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func ReplayTransaction(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"replay-transaction",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func ReplayBatch(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"replay-batch",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func ReplayCheckpoint(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"replay-checkpoint",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func Faucet(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"faucet",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func ExecuteSignedTx(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"execute-signed-tx",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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

func ExecuteCombinedSignedTx(basePath string) string {
	cmd := exec.Command(
		basePath,
		"client",
		"execute-combined-signed-tx",
	)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
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
