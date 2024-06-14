package spec

import (
	"gorm.io/gorm"
)

type EvmDepositInfo struct {
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
	State       string  `json:"state" gorm:"not null;default:'pending'"`
	L2TxHash    string  `gorm:"not null;"`
}

type EvmWithdrawInfo struct {
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
	State       string  `json:"state" gorm:"not null;default:'pending'"`
	L1TxHash    string  `gorm:"not null;"`
}
