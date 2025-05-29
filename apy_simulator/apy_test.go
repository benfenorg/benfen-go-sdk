package main

import (
	"fmt"
	"testing"
)

func Test_Apy(t *testing.T) {
	index := 7299
	nft_generate_bfc := OutputNftBfcAmount(index)
	fmt.Printf("the nft_generate_bfc is %.2f", nft_generate_bfc)

	index = 7300
	nft_generate_bfc = OutputNftBfcAmount(index)
	fmt.Printf("the nft_generate_bfc is %.2f", nft_generate_bfc)

}
