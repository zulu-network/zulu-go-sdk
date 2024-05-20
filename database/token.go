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

func (db *Database) ListZuluTokenInfoByType(tokenType int) (*[]spec.ZuluTokenInfo, error) {
	var tokenInfos []spec.ZuluTokenInfo
	if err := db.DB.Where("token_type = ?", tokenType).Find(&tokenInfos).Error; err != nil {
		return nil, err
	}
	return &tokenInfos, nil
}

func (db *Database) ListZuluTokenInfo() (*[]spec.ZuluTokenInfo, error) {
	var tokenInfos []spec.ZuluTokenInfo
	if err := db.DB.Find(&tokenInfos).Error; err != nil {
		return nil, err
	}
	return &tokenInfos, nil
}

func (db *Database) CreateL1TokenInfo(l1ti *spec.L1TokenInfo) error {
	return db.DB.Create(l1ti).Error
}

func (db *Database) GetL1TokenInfo(name string) (*spec.L1TokenInfo, error) {
	var l1TokenInfo spec.L1TokenInfo
	if err := db.DB.Where("name = ?", name).First(&l1TokenInfo).Error; err != nil {
		return nil, err
	}
	return &l1TokenInfo, nil
}
