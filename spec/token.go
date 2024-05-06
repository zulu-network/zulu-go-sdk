package spec

import (
	"gorm.io/gorm"
)

type ZuluTokenInfo struct {
	gorm.Model
	Address   string `gorm:"unique;not null;"`
	Hash      string `gorm:"not null;"`
	Name      string `gorm:"not null;"`
	Symbol    string `gorm:"not null;"`
	Decimals  uint8  `gorm:"not null;"`
	Cap       string `gorm:"not null;"`
	TokenType string `gorm:"not null;"`
}
