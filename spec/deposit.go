package spec

import (
	"gorm.io/gorm"
)

const (
	DepositTxStatePending    = "pending"
	DepositTxStateProcessing = "processing"
	DepositTxStateFailed     = "failed"
	DepositTxStateDone       = "done"
)

const (
	DepositTypeBTC   = "btc"
	DepositTypeBrc20 = "brc20"
	DepositTypeRunes = "runes"
	DepositTypeEvm   = "evm"
)

type ZuluDepositRecord struct {
	gorm.Model
	FromTxHash  string `gorm:"unique;not null;"`
	Coin        string `gorm:"not null;"`
	FromAddress string `gorm:"not null;"`
	ToAddress   string `gorm:"not null;"`
	Amount      string `gorm:"not null;"`
	Decimals    int    `gorm:"not null;"`
	Type        string `gorm:"not null;"`
}

type ZuluDepositInfo struct {
	gorm.Model
	CoboID      string `gorm:"unique"`
	Coin        string
	ChainCode   string
	DisplayCode string
	FromAddress string
	ToAddress   string
	FromTxHash  string `gorm:"unique"`
	ToTxHash    string
	Amount      string `gorm:"type:numeric"`
	AbsAmount   string `gorm:"type:numeric"`
	BridgeFee   string `gorm:"type:numeric"`
	Decimals    int
	Type        string
	BlockHeight int
	State       string `gorm:"not null;default:'pending'"`
}

// Deprecated
type ZuluDepositTxInfo struct {
	gorm.Model
	TransactionID string `gorm:"unique;not null;"`
	L1Address     string `gorm:"not null;"`
	L2Address     string `gorm:"not null;"`
	Tick          string `gorm:"not null;"`
	Type          string `gorm:"not null;"`
	Decimals      int    `gorm:"not null;"`
	Amount        string `gorm:"not null;"`
	L2TxHash      string `gorm:"not null;"`
	State         string `gorm:"not null;default:'pending'"`
}
