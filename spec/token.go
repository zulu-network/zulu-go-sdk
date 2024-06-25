package spec

import (
	"gorm.io/gorm"
)

const (
	TokenTypeBTC   = 0
	TokenTypeBrc20 = 1
	TokenTypeRunes = 2

	TokenTypeETH  = 1000
	TokenTypeUSDT = 1001
	TokenTypeUSDC = 1002
)

type ZuluTokenInfo struct {
	gorm.Model
	Address   string `gorm:"unique;not null;"`
	Hash      string `gorm:"not null;"`
	Name      string `gorm:"not null;"`
	Symbol    string `gorm:"not null;"`
	Decimals  uint8  `gorm:"not null;"`
	Cap       string `gorm:"not null;"`
	TokenType int    `gorm:"not null;"`
}

type L1TokenInfo struct {
	gorm.Model
	Name      string `gorm:"unique;not null;"`
	Symbol    string `gorm:"not null;"`
	RunesID   string `gorm:"not null;"`
	Type      int    `gorm:"not null;"`
	Decimal   int    `gorm:"not null;"`
	MaxSupply string `gorm:"not null;"`
}
