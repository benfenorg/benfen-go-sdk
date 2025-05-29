package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

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
