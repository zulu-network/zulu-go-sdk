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
	FromTokenName      string `gorm:"unique;not null;"`
	FromChainCode      string `gorm:"not null"`
	FromChainID        string `gorm:"not null;index:idx_tokenaddress_chainid,unique"`
	FromTokenDecimals  uint8  `gorm:"not null"`
	FromTokenAddress   string `gorm:"not null"`
	RecipientAddress   string `gorm:"not null"`
	DestTokenName      string `gorm:"not null;"`
	DestTokenSymbol    string `gorm:"not null;"`
	DestTokenDecimals  uint8  `gorm:"not null;"`
	DestTokenMaxSupply string `gorm:"not null;"`
	DestTokenType      int    `gorm:"not null;"`
	DestTokenAddress   string `gorm:"not null;index:idx_tokenaddress_chainid,unique"`
	DestTokenHash      string `gorm:"not null;"`
	BridgeFee          string `gorm:"not null;"`
	MinimumAmount      string `gorm:"not null;"`
	LogoUrl            string `gorm:"not null;"`
}

type CrossTokenInfo struct {
	gorm.Model
	Name             string `gorm:"unique;not null;"`
	Symbol           string `gorm:"not null;"`
	Decimals         uint8  `gorm:"not null;"`
	ChainID          string `gorm:"not null;"`
	ChainCode        string `gorm:"not null;"`
	TokenAddress     string `gorm:"unique;not null"`
	RecipientAddress string `gorm:"not null"`
	ServiceFee       string `gorm:"not null;"`
	MinimumAmount    string `gorm:"not null;"`
	LogoUrl          string `gorm:"not null;"`
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
