package database

import (
	"github.com/zulu-network/zulu-go-sdk/spec"
)

func (db *Database) CreateZuluTokenInfo(ti *spec.ZuluTokenInfo) error {
	return db.DB.Create(ti).Error
}

func (db *Database) GetZuluTokenInfo(name, symbol string, decimals uint8, maxSupply string, tokenType int) (*spec.ZuluTokenInfo, error) {
	var tokenInfo spec.ZuluTokenInfo
	if err := db.DB.Where("dest_token_name = ? AND dest_token_symbol = ? AND dest_token_decimals = ? AND dest_token_max_supply = ? AND dest_token_type = ?", name, symbol, decimals, maxSupply, tokenType).First(&tokenInfo).Error; err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

func (db *Database) GetZuluTokenInfoByAddressAndChainID(address, chainID string) (*spec.ZuluTokenInfo, error) {
	var tokenInfo spec.ZuluTokenInfo
	if err := db.DB.Where("dest_token_address = ? AND from_chain_id = ?", address, chainID).First(&tokenInfo).Error; err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

func (db *Database) GetZuluTokenInfoByFromTokenName(fromTokenName string) (*spec.ZuluTokenInfo, error) {
	var tokenInfo spec.ZuluTokenInfo
	if err := db.DB.Where("from_token_name = ?", fromTokenName).First(&tokenInfo).Error; err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

func (db *Database) UpdateZuluTokenInfo(tx *spec.ZuluTokenInfo) error {
	return db.DB.Save(tx).Error
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

// Deprecated
func (db *Database) CreateL1TokenInfo(l1ti *spec.L1TokenInfo) error {
	return db.DB.Create(l1ti).Error
}

// Deprecated
func (db *Database) GetL1TokenInfo(name string) (*spec.L1TokenInfo, error) {
	var l1TokenInfo spec.L1TokenInfo
	if err := db.DB.Where("name = ?", name).First(&l1TokenInfo).Error; err != nil {
		return nil, err
	}
	return &l1TokenInfo, nil
}

// Deprecated
func (db *Database) ListL1TokenInfo() (*[]spec.L1TokenInfo, error) {
	var tokenInfos []spec.L1TokenInfo
	if err := db.DB.Find(&tokenInfos).Error; err != nil {
		return nil, err
	}
	return &tokenInfos, nil
}
