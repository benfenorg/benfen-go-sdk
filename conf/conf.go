package conf

import (
	"github.com/benfenorg/benfen-go-sdk/benfen-go-sdk/move_types"
	"github.com/coming-chat/go-sui/v2/sui_types"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chains        []*Data_Chain  `protobuf:"bytes,1,rep,name=chains,proto3" json:"chains,omitempty"`
	Database      *Data_Database `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	Redis         *Data_Redis    `protobuf:"bytes,3,opt,name=redis,proto3" json:"redis,omitempty"`
	BlockCrawlRpc string         `protobuf:"bytes,4,opt,name=block_crawl_rpc,json=blockCrawlRpc,proto3" json:"block_crawl_rpc,omitempty"`
	Admin         *Data_Admin    `protobuf:"bytes,5,opt,name=admin,proto3" json:"admin,omitempty"`
}

type Data_Chain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Benfen, BenfenTEST
	Rpc            string   `protobuf:"bytes,2,opt,name=rpc,proto3" json:"rpc,omitempty"`
	DexPackageId   string   `protobuf:"bytes,3,opt,name=dex_package_id,json=dexPackageId,proto3" json:"dex_package_id,omitempty"`
	DexStartedTime int64    `protobuf:"varint,4,opt,name=dex_started_time,json=dexStartedTime,proto3" json:"dex_started_time,omitempty"`
	StableCoins    []string `protobuf:"bytes,5,rep,name=stable_coins,json=stableCoins,proto3" json:"stable_coins,omitempty"`
	GlobalConfigId string   `protobuf:"bytes,6,opt,name=global_config_id,json=globalConfigId,proto3" json:"global_config_id,omitempty"`
}

type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver  string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source  string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	IsDebug bool   `protobuf:"varint,3,opt,name=is_debug,json=isDebug,proto3" json:"is_debug,omitempty"`
}

type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Db           int32                `protobuf:"varint,4,opt,name=db,proto3" json:"db,omitempty"`
	DialTimeout  *durationpb.Duration `protobuf:"bytes,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,6,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	PoolSize     int32                `protobuf:"varint,8,opt,name=pool_size,json=poolSize,proto3" json:"pool_size,omitempty"`
	MinIdleConns int32                `protobuf:"varint,9,opt,name=min_idle_conns,json=minIdleConns,proto3" json:"min_idle_conns,omitempty"`
	MaxRetries   int32                `protobuf:"varint,10,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
}

type Data_Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet string `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
}

type VaultInfo struct {
	VaultID          sui_types.SuiAddress  `json:"vault_id"`
	PositionNumber   uint32                `json:"position_number"`
	State            uint8                 `json:"state"`
	StateCounter     uint32                `json:"state_counter"`
	MaxCounterTimes  uint32                `json:"max_counter_times"`
	LastSqrtPrice    move_types.U128       `json:"last_sqrt_price"`
	CoinABalance     uint64                `json:"coin_a_balance"`
	CoinBBalance     uint64                `json:"coin_b_balance"`
	CoinAType        move_types.Identifier `json:"coin_a_type"`
	CoinBType        move_types.Identifier `json:"coin_b_type"`
	TickSpacing      uint32                `json:"tick_spacing"`
	SpacingTimes     uint32                `json:"spacing_times"`
	Liquidity        move_types.U128       `json:"liquidity"`
	CurrentSqrtPrice move_types.U128       `json:"current_sqrt_price"`
	CurrentTickIndex uint32                `json:"current_tick_index"`
	IsPause          bool                  `json:"is_pause"`
	Index            uint64                `json:"index"`
	BasePoint        uint64                `json:"base_point"`
}

type VaultDisplay struct {
	CoinABalance uint64  `json:"coin_a_balance"`
	CoinBBalance uint64  `json:"coin_b_balance"`
	Price        float64 `json:"price"`
}
