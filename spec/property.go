package spec

import (
	"gorm.io/gorm"
)

type ZuluProperty struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Value string `gorm:"not null"`
}
