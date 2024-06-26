package spec

import (
	"gorm.io/gorm"
)

type Platform struct {
	Name         string
	Symbol       string
	TokenAddress string
}

type CoinMarketInfo struct {
	gorm.Model
	Name   string  `gorm:"not null;"`
	Symbol string  `gorm:"not null;unique;"`
	Price  float64 `gorm:"not null;"`
	Platform
}

type CoinBalanceInfo struct {
	gorm.Model
	Address     string `gorm:"not null;"`
	Coin        string `gorm:"not null;unique;"`
	ChainCode   string `gorm:"not null;"`
	DisplayCode string `gorm:"not null;"`
	Description string `gorm:"not null;"`
	Balance     string `gorm:"not null;"`
	Decimal     int    `gorm:"not null;"`
}

type CoinAmount struct {
	Coin        string  `json:"coin"`
	ChainCode   string  `json:"chain_code"`
	DisplayCode string  `json:"display_code"`
	Amount      uint64  `json:"amount"`
	AbsAmount   float64 `json:"abs_amount"`
	Decimals    int     `json:"decimals"`
}
