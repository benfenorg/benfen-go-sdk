package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	"time"
)

type AirdropEvent struct {
	Sender bfc_types.BfcAddress `json:"sender"`
	// Move event type.
	Type string `json:"type"`
	// Parsed json value of the event
	ParsedJson interface{} `json:"parsedJson,omitempty"`
	// Base 58 encoded bcs bytes of the move event
	Bcs string `json:"bcs"`
}

func PrintJson(data any) {
	body, _ := json.Marshal(data)
	var str bytes.Buffer
	_ = json.Indent(&str, body, "", "    ")
	fmt.Println(str.String())
}

var (
	M1Mnemonic = "xxxxxx"

	// default address
	//Address, _ = bfc_types.NewAddressFromHex("0x7419050e564485685f306e20060472fca1b3a4453b41bdace0010624801b11ea")
	Address, _ = bfc_types.NewAddressFromHex("BFC7419050e564485685f306e20060472fca1b3a4453b41bdace0010624801b11ea6cd4")
)

//func M1Address(t *testing.T) *BfcAddress {
//	return Address
//}

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

type BfcAddress = bfc_types.BfcAddress

func ConvertBfcAddressToHexAddress(address string) string {

	length := len(address)
	address = address[3 : length-4]

	address = "0x" + address
	return address
}

func ConvertBfcAddressToHexAddressWithCheck(address string) string {
	if strings.HasPrefix(address, "0x") {
		return address
	}
	length := len(address)
	address = address[3 : length-4]

	address = "0x" + address
	return address
}

func GetTimeStamp() uint64 {
	timestamp := time.Now().Unix()
	return uint64(timestamp)
}
