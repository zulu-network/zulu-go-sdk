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

type StakeDepositRecord struct {
	gorm.Model
	Coin        string  `json:"coin" gorm:"not null"`
	ChainCode   string  `json:"chain_code" gorm:"not null"`
	DisplayCode string  `json:"display_code" gorm:"not null"`
	Decimals    int     `json:"decimals" gorm:"not null"`
	Amount      uint64  `json:"amount" gorm:"not null"`
	AbsAmount   float64 `json:"abs_amount" gorm:"not null"`
	FromAddress string  `json:"from_address" gorm:"not null"`
	ToAddress   string  `json:"to_address" gorm:"not null"`
	TxHash      string  `json:"tx_hash" gorm:"unique,not null"`
}

type CoboDepositInfo struct {
	gorm.Model
	CoboID      string  `json:"cobo_id" gorm:"unique;not null"`
	Status      int     `json:"status" gorm:"not null"`
	Coin        string  `json:"coin" gorm:"not null"`
	ChainCode   string  `json:"chain_code" gorm:"not null"`
	DisplayCode string  `json:"display_code" gorm:"not null"`
	Decimals    int     `json:"decimals" gorm:"not null"`
	Amount      uint64  `json:"amount" gorm:"not null"`
	AbsAmount   float64 `json:"abs_amount" gorm:"not null"`
	FromAddress string  `json:"from_address" gorm:"not null"`
	ToAddress   string  `json:"to_address" gorm:"not null"`
	TxHash      string  `json:"tx_hash" gorm:"not null"`
	TxType      int     `json:"tx_type" gorm:"not null"`
	BlockHeight int     `json:"block_height" gorm:"not null"`
}

type CoboWithdrawInfo struct {
	gorm.Model
	CoboID      string  `json:"cobo_id" gorm:"unique;not null"`
	Status      int     `json:"status" gorm:"not null"`
	Coin        string  `json:"coin" gorm:"not null"`
	ChainCode   string  `json:"chain_code" gorm:"not null"`
	DisplayCode string  `json:"display_code" gorm:"not null"`
	Decimals    int     `json:"decimals" gorm:"not null"`
	Amount      uint64  `json:"amount" gorm:"not null"`
	AbsAmount   float64 `json:"abs_amount" gorm:"not null"`
	FromAddress string  `json:"from_address" gorm:"not null"`
	ToAddress   string  `json:"to_address" gorm:"not null"`
	TxHash      string  `json:"tx_hash" gorm:"not null"`
	TxType      int     `json:"tx_type" gorm:"not null"`
	BlockHeight int     `json:"block_height" gorm:"not null"`
}

type CoinAmount struct {
	Coin        string  `json:"coin"`
	ChainCode   string  `json:"chain_code"`
	DisplayCode string  `json:"display_code"`
	Amount      uint64  `json:"amount"`
	AbsAmount   float64 `json:"abs_amount"`
	Decimals    int     `json:"decimals"`
}
