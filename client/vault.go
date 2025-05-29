package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/bfc_types"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/conf"
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/types"
	"github.com/coming-chat/go-sui/v2/sui_types"
	"github.com/fardream/go-bcs/bcs"
	"github.com/go-kratos/kratos/v2/log"
	"math"
	"math/big"
	"os"
)

const NativeCoin = "0x2::bfc::BFC"
const ChainName = "test"

type suiAddress = sui_types.SuiAddress

type BenfenChainClient struct {
	cli            *Client
	stableCoins    []string
	nativeCoin     string
	dexPackageId   string
	dexStartedTime int64
}

// NewBenfenChainService creates a new BenfenChainService instance of type `BenfenChainService“
func NewBenfenChainService(c *conf.Data, logger log.Logger) *BenfenChainService {
	svc := &BenfenChainService{
		logHelper: log.NewHelper(log.With(logger, "module", "data/benfen_rpc")),
		clients:   make(map[string]*BenfenChainClient),
	}
	for _, chain := range c.Chains {
		cli, err := Dial(chain.Rpc)
		if err != nil {
			svc.logHelper.Fatalf("BenfenRpcEndpoint dial failed, %v", err)
		}
		benfenCli := &BenfenChainClient{
			cli:            cli,
			stableCoins:    chain.StableCoins,
			nativeCoin:     NativeCoin,
			dexPackageId:   chain.DexPackageId,
			dexStartedTime: chain.DexStartedTime,
		}
		svc.clients[chain.Name] = benfenCli
	}
	return svc
}

type BenfenChainService struct {
	logHelper *log.Helper
	clients   map[string]*BenfenChainClient
}

func GetTestBenfenRpc(rpc string) *BenfenChainService {
	if rpc == "" {
		rpc = "https://testrpc.benfen.org"
	}

	return NewBenfenChainService(
		&conf.Data{
			Chains: []*conf.Data_Chain{
				{
					Name:        ChainName,
					Rpc:         rpc,
					StableCoins: []string{},
				},
			},
		},
		log.NewStdLogger(os.Stdout),
	)
}

// 获取稳定币池子信息
func (bfSvc *BenfenChainService) VaultInfo(ctx context.Context, chainName, stableCoin string) (
	*conf.VaultInfo,
	error,
) {
	fromAddr, err := bfc_types.NewAddressFromHex("0x0")
	if err != nil {
		bfSvc.logHelper.WithContext(ctx).Errorf("Failed to new address, %v", err)
		return nil, err
	}
	bfCli, ok := bfSvc.clients[chainName]
	if !ok {
		return nil, errors.New("chain name not found")
	}
	resp, err := bfCli.cli.DevInspectTransactionBlockV2(
		context.Background(),
		*fromAddr,
		"0xc8::bfc_system::vault_info",
		[]string{
			stableCoin,
		},
		[]*DevInspectArgs{
			{
				Type:    "object",
				Value:   "0xc9",
				Version: 1,
			},
		},
	)
	if err != nil {
		bfSvc.logHelper.WithContext(ctx).Errorf("Failed to get vault_info by bfSvc, %v", err)
		return nil, err
	}
	var vaultInfo conf.VaultInfo
	if err := unmarshalDevInspectResp(resp, &vaultInfo); err != nil {
		bfSvc.logHelper.WithContext(ctx).Errorf("Failed to unmarshal resp: %v", err)
		return nil, err
	}
	return &vaultInfo, nil
}

// 这里取 ReturnValues 中的第一个，如果返回是一个 struct 则使用这个方法
func unmarshalDevInspectResp(results *types.DevInspectResults, dst any) error {
	//fmt.Println(len(results.Results))
	return unmarshalDevInspectValue(results.Results[0].ReturnValues[0], dst)
}

// 如果ReturnValues 返回的是一个数组，需要对每一个进行 unmarshal
func unmarshalDevInspectValue(value types.ReturnValueType, dst any) error {
	resBytes := []byte{}
	interfaceData := value.([]interface{})[0]
	switch interfaceData.(type) {
	case []interface{}:
		for _, item := range interfaceData.([]interface{}) {
			resBytes = append(resBytes, uint8(item.(float64)))
		}
	}
	if _, err := bcs.Unmarshal(resBytes, dst); err != nil {
		return err
	}
	return nil
}

func SqrtPrice2Price(sqrtPriceX64Str string) (float64, error) {
	sqrtPriceX64, ok := new(big.Float).SetString(sqrtPriceX64Str)
	if !ok {
		return 0, fmt.Errorf("SqrtPrice2Price.format %v error", sqrtPriceX64Str)
	}

	sqrtPrice := big.NewFloat(0).Quo(sqrtPriceX64, big.NewFloat(math.Pow(2, 64)))
	ret := big.NewFloat(0).Mul(sqrtPrice, sqrtPrice)
	price, _ := ret.Float64()
	return price, nil
}
