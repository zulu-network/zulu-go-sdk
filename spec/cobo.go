package spec

import (
	"gorm.io/gorm"
)

type CoboDepositInfo struct {
	gorm.Model
	CoboID      string `json:"cobo_id" gorm:"unique;not null"`
	Status      int    `json:"status" gorm:"not null"`
	Coin        string `json:"coin" gorm:"not null"`
	Decimals    int    `json:"decimals" gorm:"not null"`
	Amount      string `json:"amount" gorm:"not null"`
	AbsAmount   string `json:"abs_amount" gorm:"not null"`
	FromAddress string `json:"from_address" gorm:"not null"`
	ToAddress   string `json:"to_address" gorm:"not null"`
	TxHash      string `json:"tx_hash" gorm:"not null"`
	TxType      int    `json:"tx_type" gorm:"not null"`
	BlockHeight int    `json:"block_height" gorm:"not null"`
}

type CoboWithdrawInfo struct {
	gorm.Model
	CoboID      string `json:"cobo_id" gorm:"unique;not null"`
	Status      int    `json:"status" gorm:"not null"`
	Coin        string `json:"coin" gorm:"not null"`
	Decimals    int    `json:"decimals" gorm:"not null"`
	Amount      string `json:"amount" gorm:"not null"`
	AbsAmount   string `json:"abs_amount" gorm:"not null"`
	FromAddress string `json:"from_address" gorm:"not null"`
	ToAddress   string `json:"to_address" gorm:"not null"`
	TxHash      string `json:"tx_hash" gorm:"not null"`
	TxType      int    `json:"tx_type" gorm:"not null"`
	BlockHeight int    `json:"block_height" gorm:"not null"`
}
