package database

import (
	"errors"

	"gorm.io/gorm"

	"github.com/zulu-network/zulu-go-sdk/spec"
)

// CoinMarket
func (db *Database) CreateCoinMarketInfo(cmi *spec.CoinMarketInfo) error {
	return db.DB.Create(cmi).Error
}

func (db *Database) CreateOrUpdateCoinMarketInfo(cmi *spec.CoinMarketInfo) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var existing spec.CoinMarketInfo
		if err := tx.Where(&spec.CoinMarketInfo{Symbol: cmi.Symbol}).First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return tx.Create(cmi).Error
			}
			return err
		}
		return tx.Model(&existing).Updates(cmi).Error
	})
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

// CoinBalance
func (db *Database) CreateCoinBalanceInfo(cbi *spec.CoinBalanceInfo) error {
	return db.DB.Create(cbi).Error
}

func (db *Database) CreateOrUpdateCoinBalanceInfo(cbi *spec.CoinBalanceInfo) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var existing spec.CoinBalanceInfo
		if err := tx.Where(&spec.CoinBalanceInfo{Coin: cbi.Coin}).First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return tx.Create(cbi).Error
			}
			return err
		}
		return tx.Model(&existing).Updates(cbi).Error
	})
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

// Staking Deposit & Withdraw
func (db *Database) CreateStakeDepositRecord(sdr *spec.StakeDepositRecord) error {
	return db.DB.Create(sdr).Error
}

func (db *Database) CreateAndGetStakeDepositRecord(tx *spec.StakeDepositRecord) (*spec.StakeDepositRecord, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.StakeDepositRecord
	if err := db.DB.Where("tx_hash = ?", tx.TxHash).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) GetStakeDepositRecord(txhash string) (*spec.StakeDepositRecord, error) {
	var sdr spec.StakeDepositRecord
	if err := db.DB.Where("tx_hash = ?", txhash).First(&sdr).Error; err != nil {
		return nil, err
	}
	return &sdr, nil
}

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
