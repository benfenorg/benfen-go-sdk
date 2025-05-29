package client

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVaultInfo(t *testing.T) {
	rpc := GetTestBenfenRpc("")
	info, err := rpc.VaultInfo(context.Background(), ChainName, "0xc8::busd::BUSD")
	require.NoError(t, err)
	// assert.Equal(t, info.BasePoint, uint64(1000000000000))
	fmt.Printf("a balance: %d\n", info.CoinABalance)
	fmt.Printf("b balance: %d\n", info.CoinBBalance)
	ans, _ := SqrtPrice2Price(info.LastSqrtPrice.String())
	fmt.Printf("price: %f\n", ans)
}
