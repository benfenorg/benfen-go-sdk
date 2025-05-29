package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func ZkLoginSignAndExecuteTx(basePath string, epoch string, net string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"zk-login-sign-and-execute-tx",
		"--max-epoch",
		epoch,
		"--network",
		net,
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

func UpdateAlias(basePath string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"zk-login-sign-and-execute-tx",
		"update-alias",
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

func list(basePath string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"list",
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

func export(basePath string, identity string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"export",
		"--key-identity",
		identity,
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

func ZkLoginInsecureSignPersonalMessage(basePath string, data string, epoch string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"zk-login-insecure-sign-personal-message",
		"--data",
		data,
		"--max-epoch",
		epoch,
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

func ZkLoginSigVerify(basePath string, bytes string, sig string, scope string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"zk-login-sig-verify",
		"--bytes",
		bytes,
		"--sig",
		sig,
		"--intent-scope",
		scope,
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

func ZkLoginEnterToken(basePath string, token string, kp string, epoch string, jwt string, key string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"zk-login-enter-token",
		"--parsed-token",
		token,
		"--kp-bigint",
		kp,
		"--max-epoch",
		epoch,
		"--jwt-randomness",
		jwt,
		"--ephemeral-key-identifier",
		key,
		"--network",
		"localnet",
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

func Unpack(basePath string, str string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"unpack",
		str,
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

func SignKms(basePath string, data string, key string, pk string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"sign-kms",
		"--data",
		data,
		"--keyid",
		key,
		"--base64pk",
		pk,
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

func Sign(basePath string, data string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"sign",
		"--data",
		data,
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

func MultiSigCombinePartialSigLegacy(
	basePath string,
	pks string,
	weights string,
	threshold string,
	sigs string,
) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"multi-sig-combine-partial-sig-legacy",
		"--pks",
		pks,
		"--weights",
		weights,
		"--threshold",
		threshold,
		"--sigs",
		sigs,
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

func MultiSigCombinePartialSig(
	basePath string,
	pks string,
	weights string,
	threshold string,
	sigs string,
) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"multi-sig-combine-partial-sig",
		"--pks",
		pks,
		"--weights",
		weights,
		"--threshold",
		threshold,
		"--sigs",
		sigs,
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

func MultiSigAddress(basePath string, pks string, weights string, threshold string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"multi-sig-address",
		"--pks",
		pks,
		"--weights",
		weights,
		"--threshold",
		threshold,
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

func DecodeMultiSig(basePath string, sig string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"decode-multi-sig",
		"--multisig",
		sig,
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

func Convert(basePath string, str string) string {
	cmd := exec.Command(
		basePath,
		"keytool",
		"convert",
		str,
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
