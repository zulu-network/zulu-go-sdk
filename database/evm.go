package database

import (
	"github.com/zulu-network/zulu-go-sdk/spec"
)

func (db *Database) CreateEvmDepositTransaction(tx *spec.EvmDepositInfo) error {
	return db.DB.Create(tx).Error
}

func (db *Database) CreateAndGetEvmDepositTransaction(tx *spec.EvmDepositInfo) (*spec.EvmDepositInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.EvmDepositInfo
	if err := db.DB.Where("cobo_id = ?", tx.CoboID).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) ListEvmDepositTransactionByAddress(fromAddress string) (*[]spec.EvmDepositInfo, error) {
	var txInfos []spec.EvmDepositInfo
	if err := db.DB.Where("from_address = ?", fromAddress).Order("created_at DESC").Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) ListUnhandledEvmDepositTransactions(number int) (*[]spec.EvmDepositInfo, error) {
	var txInfos []spec.EvmDepositInfo
	if err := db.DB.Order("created_at asc").Limit(number).Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) GetEvmAmountsByFromAddress(fromAddress string) ([]spec.CoinAmount, error) {
	var coinAmounts []spec.CoinAmount
	err := db.DB.Model(&spec.EvmDepositInfo{}).
		Select("coin, chain_code, display_code, decimals, COALESCE(SUM(amount), 0) as amount, COALESCE(SUM(abs_amount), 0) as abs_amount").
		Where("from_address = ?", fromAddress).
		Group("coin, chain_code, display_code, decimals, chain_code"). // 添加 chain_code 到 GROUP BY 子句中
		Scan(&coinAmounts).Error

	return coinAmounts, err
}

func (db *Database) CreateEvmWithdrawTransaction(tx *spec.EvmWithdrawInfo) error {
	return db.DB.Create(tx).Error
}

func (db *Database) CreateAndGetEvmWithdrawTransaction(tx *spec.EvmWithdrawInfo) (*spec.EvmWithdrawInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.EvmWithdrawInfo
	if err := db.DB.Where("cobo_id = ?", tx.CoboID).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) ListEvmWithdrawTransactionByAddress(toAddress string) (*[]spec.EvmWithdrawInfo, error) {
	var txInfos []spec.EvmWithdrawInfo
	if err := db.DB.Where("to_address = ?", toAddress).Order("created_at DESC").Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}
