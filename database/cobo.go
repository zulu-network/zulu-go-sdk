package database

import (
	"github.com/zulu-network/zulu-go-sdk/spec"
)

func (db *Database) CreateCoboDepositTransaction(tx *spec.CoboDepositInfo) error {
	return db.DB.Create(tx).Error
}

func (db *Database) CreateAndGetCoboDepositTransaction(tx *spec.CoboDepositInfo) (*spec.CoboDepositInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.CoboDepositInfo
	if err := db.DB.Where("cobo_id = ?", tx.CoboID).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) GetAmountsByFromAddress(fromAddress string) ([]spec.CoinAmount, error) {
	var coinAmounts []spec.CoinAmount
	err := db.DB.Model(&spec.CoboDepositInfo{}).
		Select("coin, chain_code, display_code, amount, abs_amount, decimals, COALESCE(SUM(amount), '0') as amount").
		Where("from_address = ?", fromAddress).
		Group("coin").
		Scan(&coinAmounts).Error

	return coinAmounts, err
}

func (db *Database) CreateCoboWithdrawTransaction(tx *spec.CoboWithdrawInfo) error {
	return db.DB.Create(tx).Error
}

func (db *Database) CreateAndGetCoboWithdrawTransaction(tx *spec.CoboWithdrawInfo) (*spec.CoboWithdrawInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.CoboWithdrawInfo
	if err := db.DB.Where("cobo_id = ?", tx.CoboID).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}
