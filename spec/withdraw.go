package spec

import (
	"gorm.io/gorm"
)

const (
	WithdrawTxStatePending    = "pending"
	WithdrawTxStateProcessing = "processing"
	WithdrawTxStateDone       = "done"
)

type ZuluWithdrawTxInfo struct {
	gorm.Model
	TxHash    string `gorm:"unique;not null;"`
	L2Address string `gorm:"not null;"`
	L1Address string `gorm:"not null;"`
	Tick      string `gorm:"not null;"`
	Amount    []byte `gorm:"not null;"`
	State     string `gorm:"not null;default:'pending'"`
}
