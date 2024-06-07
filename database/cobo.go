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

func (db *Database) ListCoboDepositTransactionByAddress(fromAddress string) (*[]spec.CoboDepositInfo, error) {
	var txInfos []spec.CoboDepositInfo
	if err := db.DB.Where("from_address = ?", fromAddress).Order("created_at DESC").Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) GetAmountsByFromAddress(fromAddress string) ([]spec.CoinAmount, error) {
	var coinAmounts []spec.CoinAmount
	err := db.DB.Model(&spec.CoboDepositInfo{}).
		Select("coin, chain_code, display_code, decimals, COALESCE(SUM(amount), 0) as amount, COALESCE(SUM(abs_amount), 0) as abs_amount").
		Where("from_address = ?", fromAddress).
		Group("coin, chain_code, display_code, decimals, chain_code"). // 添加 chain_code 到 GROUP BY 子句中
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

func (db *Database) ListCoboWithdrawTransactionByAddress(toAddress string) (*[]spec.CoboWithdrawInfo, error) {
	var txInfos []spec.CoboWithdrawInfo
	if err := db.DB.Where("to_address = ?", toAddress).Order("created_at DESC").Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}
