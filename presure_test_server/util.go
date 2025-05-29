package main

import (
	"strconv"
	"testing"

	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"github.com/shopspring/decimal"

	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/account"
	"github.com/stretchr/testify/require"
)

var (
	M1Mnemonic = "monkey tragic drive owner fade mimic taxi despair endorse peasant amused woman"

	Address, _ = bfc_types.NewAddressFromHex("0x7419050e564485685f306e20060472fca1b3a4453b41bdace0010624801b11ea")
)

func M1Account(t *testing.T) *account.Account {
	a, err := account.NewAccountWithMnemonic(M1Mnemonic)
	require.NoError(t, err)
	return a
}

//func M1Address(t *testing.T) *BfcAddress {
//	return Address
//}

func Signer(t *testing.T) *account.Account {
	return M1Account(t)
}

type BFC float64

func (s BFC) Int64() int64 {
	return int64(s * 1e9)
}
func (s BFC) Uint64() uint64 {
	return uint64(s * 1e9)
}
func (s BFC) Decimal() decimal.Decimal {
	return decimal.NewFromInt(s.Int64())
}
func (s BFC) String() string {
	return strconv.FormatInt(s.Int64(), 10)
}
