package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"strings"
	"syscall"
)

func aes128Encrypt(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 对明文进行填充
	blockSize := block.BlockSize()
	padding := blockSize - len(plaintext)%blockSize
	paddedPlaintext := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// 创建AES加密器
	mode := cipher.NewCBCEncrypter(block, key[:blockSize])
	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	return ciphertext, nil
}

func aes128Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建AES解密器
	blockSize := block.BlockSize()

	mode := cipher.NewCBCDecrypter(block, key[:blockSize])
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除填充数据
	padding := int(plaintext[len(plaintext)-1])
	plaintext = plaintext[:len(plaintext)-padding]

	return plaintext, nil
}

func main() {
	//key := []byte("0123456789abcdef") // 16字节的密钥

	fmt.Print("请输入加密说使用的私钥【1-16位】: ")
	var input string
	bytePassword, _ := terminal.ReadPassword(syscall.Stdin)
	input = string(bytePassword)

	if len(input) > 16 {
		fmt.Print("用户输入密码必须【1-16位】")
		return
	}
	if len(input) < 16 {
		padding := 16 - len(input)
		input = input + strings.Repeat("a", padding)
	}

	plaintext := []byte("Hello, AES-128!") // 要加密的明文

	keystore := ReadFile()
	if len(keystore) > 0 {
		plaintext = []byte(keystore)
	}

	ciphertext, err := aes128Encrypt([]byte(input), plaintext) // 加密

	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}

	// 使用base64编码输出加密后的结果
	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("加密结果:", encodedCiphertext)

	Write_file([]byte(encodedCiphertext), "bfc.keystore_crypted")

	fmt.Print("请输入解密使用的私钥【1-16位】: ")
	var userKey string
	bytePassword, _ = terminal.ReadPassword(syscall.Stdin)
	userKey = string(bytePassword)
	if len(userKey) > 16 {
		fmt.Print("用户输入密码必须【1-16位】")
		return
	}
	if len(userKey) < 16 {
		padding := 16 - len(userKey)
		userKey = userKey + strings.Repeat("a", padding)
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		fmt.Println("解码失败:", err)
		return
	}

	decryptedPlaintext, err := aes128Decrypt([]byte(userKey), decodedCiphertext) // 解密
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}

	fmt.Println("解密结果:", string(decryptedPlaintext))
	Write_file(decryptedPlaintext, "bfc.keystore_decrypted")
}
