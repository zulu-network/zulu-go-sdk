package spec

import (
	"gorm.io/gorm"
)

const (
	DepositTxStatePending    = "pending"
	DepositTxStateProcessing = "processing"
	DepositTxStateDone       = "done"
)

type ZuluDepositTxInfo struct {
	gorm.Model
	TransactionID string `gorm:"unique;not null;"`
	L1Address     string `gorm:"not null;"`
	L2Address     string `gorm:"not null;"`
	Tick          string `gorm:"not null;"`
	Amount        string `gorm:"not null;"`
	State         string `gorm:"not null;default:'pending'"`
}
