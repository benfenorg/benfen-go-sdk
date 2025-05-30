package bfc_types

import "github.com/benfenorg/benfen-go-sdk/move_types"

const (
	StakingPoolModuleName = move_types.Identifier("staking_pool")
	StakedBfcStructName   = move_types.Identifier("StakedBfc")

	AddStakeMulCoinFunName = move_types.Identifier("request_add_stake_mul_coin")
	AddStakeFunName        = move_types.Identifier("request_add_stake")
	WithdrawStakeFunName   = move_types.Identifier("request_withdraw_stake")
)
