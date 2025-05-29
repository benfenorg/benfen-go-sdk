package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func MakeValidatorInfo(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"make-validator-info",
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

func BecomeCandidate(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"become-candidate",
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

func JoinCommittee(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"join-committee",
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

func LeaveCommittee(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"leave-committee",
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

func DisplayMetadata(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"display-metadata",
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

func UpdateMetadata(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"update-metadata",
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

func UpdateGasPrice(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"update-gas-price",
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

func ReportValidator(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"report-validator",
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

func SerializePayloadPop(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"serialize-payload-pop",
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

func DisplayGasPriceUpdateRawTx(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"display-gas-price-update-raw-txn",
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

func InitAdminCapability(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"init-admin-capability",
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

func AddAdminCapability(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"add-admin-capability",
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

func RemoveAdminCapability(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"remove-admin-capability",
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

func AddOperationCapability(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"add-operation-capability",
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

func RemoveOperationCapability(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"remove-operation-capability",
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

func SetOraclePriceAddress(basePath string) string {
	cmd := exec.Command(
		basePath,
		"validator",
		"set-oracle-price-address",
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
