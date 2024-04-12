package spec

import (
	"gorm.io/gorm"
)

type BtcDepositTxInfo struct {
	gorm.Model
	TransactionID string `gorm:"unique;not null;"`
	L1Address     string `gorm:"not null;"`
	L2Address     string `gorm:"not null;"`
	Tick          string `gorm:"not null;"`
	Amount        []byte `gorm:"not null;"`
	State         string `gorm:"not null;default:'pending'"`
}

type Brc20DepositTxInfo struct {
	gorm.Model
	InscriptionID string `gorm:"unique;not null;"`
	L1Address     string `gorm:"not null;"`
	L2Address     string `gorm:"not null;"`
	Tick          string `gorm:"not null;"`
	Amount        []byte `gorm:"not null;"`
	State         string `gorm:"not null;default:'pending'"`
}
