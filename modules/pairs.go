package modules

import "github.com/shopspring/decimal"

type Pairs struct {
	Id       uint64          `json:"id"`
	Token0Id uint64          `json:"token0Id"`
	Token1Id uint64          `json:"token1Id"`
	Reserve0 decimal.Decimal `json:"reserve0"`
	Reserve1 decimal.Decimal `json:"reserve1"`
}
