package spec

import (
	"math/big"

	"gorm.io/gorm"
)

const (
	TokenTypeBTC   = 0
	TokenTypeBrc20 = 1
	TokenTypeRunes = 2
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
	Name      string   `json:"name"`
	Symbol    string   `json:"symbol"`
	RunesID   string   `json:"runes_id"`
	Type      int      `json:"type"`
	Decimal   int      `json:"decimal"`
	MaxSupply *big.Int `json:"maxSupply"`
}
