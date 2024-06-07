package database

import (
	"github.com/zulu-network/zulu-go-sdk/spec"
)

func (db *Database) CreateCoinMarketInfo(cmi *spec.CoinMarketInfo) error {
	return db.DB.Create(cmi).Error
}

func (db *Database) CreateOrUpdateCoinMarketInfo(cmi *spec.CoinMarketInfo) error {
	return db.DB.Save(cmi).Error
}

func (db *Database) GetCoinMarketInfo(name string) (*spec.CoinMarketInfo, error) {
	var coin spec.CoinMarketInfo
	if err := db.DB.Where("name = ?", name).First(&coin).Error; err != nil {
		return nil, err
	}
	return &coin, nil
}

func (db *Database) GetCoinMarketInfoBySymbol(symbol string) (*spec.CoinMarketInfo, error) {
	var coin spec.CoinMarketInfo
	if err := db.DB.Where("symbol = ?", symbol).First(&coin).Error; err != nil {
		return nil, err
	}
	return &coin, nil
}

func (db *Database) CreateCoinBalanceInfo(cbi *spec.CoinBalanceInfo) error {
	return db.DB.Create(cbi).Error
}

func (db *Database) CreateOrUpdateCoinBalanceInfo(cbi *spec.CoinBalanceInfo) error {
	return db.DB.Save(cbi).Error
}

func (db *Database) GetCoinBalanceInfo(address, coin string) (*spec.CoinBalanceInfo, error) {
	var cbi spec.CoinBalanceInfo
	if err := db.DB.Where("address = ? AND coin = ?", address, coin).First(&cbi).Error; err != nil {
		return nil, err
	}
	return &cbi, nil
}

func (db *Database) GetCoinBalanceInfoByDisplayCode(address, displayCode string) (*spec.CoinBalanceInfo, error) {
	var cbi spec.CoinBalanceInfo
	if err := db.DB.Where("address = ? AND display_code = ?", address, displayCode).First(&cbi).Error; err != nil {
		return nil, err
	}
	return &cbi, nil
}
