package spec

import (
	"gorm.io/gorm"
)

const (
	WithdrawTxStatePending       = "pending"
	WithdrawTxStateProcessing    = "processing"
	WithdrawTxStateProcessFailed = "failed"
	WithdrawTxStateDone          = "done"
)

type ZuluWithdrawInfo struct {
	gorm.Model
	CoboID      *string `gorm:"unique"`
	Coin        string
	FromAddress CITEXT `gorm:"type:citext"`
	ToAddress   CITEXT `gorm:"type:citext"`
	DestChain   string
	FromTxHash  string `gorm:"unique"`
	ToTxHash    string
	Amount      string
	AbsAmount   string
	Type        int
	Decimals    int
	BlockHeight int
	State       string `gorm:"not null;default:'pending'"`
}

// Deprecated
type ZuluWithdrawTxInfo struct {
	gorm.Model
	TxHash    string `gorm:"unique;not null;"`
	L2Address string `gorm:"not null;"`
	L1Address string `gorm:"not null;"`
	Tick      string `gorm:"not null;"`
	Type      int    `gorm:"not null;"`
	Decimals  int    `gorm:"not null;"`
	Amount    string `gorm:"not null;"`
	L1TxHash  string `gorm:"not null;"`
	State     string `gorm:"not null;default:'pending'"`
}
