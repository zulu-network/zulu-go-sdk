package database

import (
	"github.com/zulu-network/zulu-go-sdk/spec"
)

func (db *Database) GetDepositTransaction(txid, fromAddress string) (*spec.ZuluDepositTxInfo, error) {
	var txInfo spec.ZuluDepositTxInfo
	if err := db.DB.Where("transaction_id = ? AND l1address = ?", txid, fromAddress).Find(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) CreateDepositTransaction(tx *spec.ZuluDepositTxInfo) error {
	if err := db.DB.Create(tx).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) CreateAndGetDepositTransaction(tx *spec.ZuluDepositTxInfo) (*spec.ZuluDepositTxInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.ZuluDepositTxInfo
	if err := db.DB.Where("transaction_id = ? AND l1address = ?", tx.TransactionID, tx.L1Address).Find(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) ListDepositTransactionByAddress(fromAddress string) (*[]spec.ZuluDepositTxInfo, error) {
	var txInfos []spec.ZuluDepositTxInfo
	if err := db.DB.Where("l1address = ?", fromAddress).Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) UpdateDepositTransaction(tx *spec.ZuluDepositTxInfo) error {
	if err := db.DB.Model(&spec.ZuluDepositTxInfo{}).Where("transaction_id = ? AND l1address = ?", tx.TransactionID, tx.L1Address).
		Update("state", tx.State).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) DeleteDepositTransaction(txid, fromAddress string) error {
	var txInfo spec.ZuluDepositTxInfo
	if err := db.DB.Where("transaction_id = ? AND l1address = ?", txid, fromAddress).Delete(&txInfo).Error; err != nil {
		return err
	}
	return nil
}

// withdraw
func (db *Database) GetZuluWithdrawTransaction(txhash, fromAddress string) (*spec.ZuluWithdrawTxInfo, error) {
	var txInfo spec.ZuluWithdrawTxInfo
	if err := db.DB.Where("txhash = ? AND l2address = ?", txhash, fromAddress).Find(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) CreateZuluWithdrawTransaction(tx *spec.ZuluWithdrawTxInfo) error {
	if err := db.DB.Create(tx).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) CreateAndGetZuluWithdrawTransaction(tx *spec.ZuluWithdrawTxInfo) (*spec.ZuluWithdrawTxInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.ZuluWithdrawTxInfo
	if err := db.DB.Where("txhash = ? AND l2address = ?", tx.TxHash, tx.L2Address).Find(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) ListZuluWithdrawTransactionByAddress(fromAddress string) (*[]spec.ZuluWithdrawTxInfo, error) {
	var txInfos []spec.ZuluWithdrawTxInfo
	if err := db.DB.Where("l2address = ?", fromAddress).Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) UpdateZuluWithdrawTransaction(tx *spec.ZuluWithdrawTxInfo) error {
	if err := db.DB.Model(&spec.ZuluWithdrawTxInfo{}).Where("txhash = ? AND l2address = ?", tx.TxHash, tx.L2Address).
		Update("state", tx.State).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) DeleteZuluWithdrawTransaction(txhash, fromAddress string) error {
	var txInfo spec.ZuluWithdrawTxInfo
	if err := db.DB.Where("txhash = ? AND l2address = ?", txhash, fromAddress).Delete(&txInfo).Error; err != nil {
		return err
	}
	return nil
}
