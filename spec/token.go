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

	TokenTypeBNB     = 2000
	TokenTypeFDUSD   = 2001
	TokenTypeBTCB    = 2002
	TokenTypeBBTC    = 2003
	TokenTypeSolvBTC = 2004
)

type ZuluTokenInfo struct {
	gorm.Model
	Name          string `gorm:"not null;index:idx_nsdmt,unique"`
	Symbol        string `gorm:"not null;index:idx_nsdmt,unique"`
	RunesID       string `gorm:"not null;"`
	Decimals      uint8  `gorm:"not null;index:idx_nsdmt,unique"`
	MaxSupply     string `gorm:"not null;index:idx_nsdmt,unique"`
	TokenType     int    `gorm:"not null;index:idx_nsdmt,unique"`
	CoboName      string `gorm:"unique;not null;"`
	ChainCode     string `gorm:"not null"`
	ChainID       string `gorm:"not null;index:idx_tokenaddress_chainid,unique"`
	TokenAddress  string `gorm:"not null;index:idx_tokenaddress_chainid,unique"`
	TokenHash     string `gorm:"not null;"`
	BridgeFee     string `gorm:"not null;"`
	MinimumAmount string `gorm:"not null;"`
	LogoUrl       string `gorm:"not null;"`
}

// Deprecated
type L1TokenInfo struct {
	gorm.Model
	Name      string `gorm:"unique;not null;"`
	Symbol    string `gorm:"not null;"`
	RunesID   string `gorm:"not null;"`
	Type      int    `gorm:"not null;"`
	BridgeFee string `gorm:"not null;"`
	Decimal   int    `gorm:"not null;"`
	MaxSupply string `gorm:"not null;"`
}
