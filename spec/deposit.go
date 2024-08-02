package spec

import (
	"gorm.io/gorm"
)

// Define a custom datatype for citext
type CITEXT string

// GormDataType defines GORM's data type for the CITEXT type
func (CITEXT) GormDataType() string {
	return "citext"
}

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
	FromAddress CITEXT `gorm:"not null;type:citext"`
	ToAddress   CITEXT `gorm:"not null;type:citext"`
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
	FromAddress CITEXT `gorm:"type:citext"`
	ToAddress   CITEXT `gorm:"type:citext"`
	FromTxHash  string
	ToTxHash    string
	Amount      string `gorm:"type:numeric"`
	AbsAmount   string `gorm:"type:numeric"`
	BridgeFee   string `gorm:"type:numeric"`
	Decimals    int
	Type        string
	BlockHeight int
	State       string `gorm:"not null;default:'pending'"`
}

type ZuluQuickCrossInfo struct {
	gorm.Model
	CoboID      string `gorm:"unique"`
	Coin        string
	ChainCode   string
	DisplayCode string
	FromAddress CITEXT `gorm:"type:citext"`
	ToAddress   CITEXT `gorm:"type:citext"`
	FromTxHash  string
	ToTxHash    string
	Amount      string `gorm:"type:numeric"`
	BtcAmount   string `gorm:"type:numeric"`
	BridgeFee   string `gorm:"type:numeric"`
	Decimals    int
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
