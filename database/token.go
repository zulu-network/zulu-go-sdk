package database

import (
	"github.com/zulu-network/zulu-go-sdk/spec"
)

func (db *Database) CreateZuluTokenInfo(ti *spec.ZuluTokenInfo) error {
	return db.DB.Create(ti).Error
}

func (db *Database) GetZuluTokenInfo(address string) (*spec.ZuluTokenInfo, error) {
	var tokenInfo spec.ZuluTokenInfo
	if err := db.DB.Where("address = ?", address).First(&tokenInfo).Error; err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

func (db *Database) GetZuluTokenInfoByName(name string) (*spec.ZuluTokenInfo, error) {
	var tokenInfo spec.ZuluTokenInfo
	if err := db.DB.Where("name = ?", name).First(&tokenInfo).Error; err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}
