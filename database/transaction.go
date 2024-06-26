package database

import (
	"github.com/zulu-network/zulu-go-sdk/spec"
)

// Deposit
func (db *Database) CreateDepositTransaction(tx *spec.ZuluDepositInfo) error {
	return db.DB.Create(tx).Error
}

func (db *Database) CreateAndGetDepositTransaction(tx *spec.ZuluDepositInfo) (*spec.ZuluDepositInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.ZuluDepositInfo
	if err := db.DB.Where("from_tx_hash = ?", tx.FromTxHash).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) ListDepositTransactionByAddress(fromAddress string) (*[]spec.ZuluDepositInfo, error) {
	var txInfos []spec.ZuluDepositInfo
	if err := db.DB.Where("from_address = ?", fromAddress).Order("created_at DESC").Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) ListUnhandledDepositTransactions(number int) (*[]spec.ZuluDepositInfo, error) {
	var txInfos []spec.ZuluDepositInfo
	if err := db.DB.Where("state = ?", spec.DepositTxStatePending).Order("created_at asc").Limit(number).Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) UpdateDepositTransaction(tx *spec.ZuluDepositInfo) error {
	return db.DB.Save(tx).Error
}

func (db *Database) CreateDepositRecord(tx *spec.ZuluDepositRecord) error {
	return db.DB.Create(tx).Error
}

func (db *Database) CreateAndGetDepositRecord(tx *spec.ZuluDepositRecord) (*spec.ZuluDepositRecord, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var rcd spec.ZuluDepositRecord
	if err := db.DB.Where("from_tx_hash = ?", tx.FromTxHash).First(&rcd).Error; err != nil {
		return nil, err
	}
	return &rcd, nil
}

func (db *Database) GetDepositRecord(txHash string) (*spec.ZuluDepositRecord, error) {
	var rcd spec.ZuluDepositRecord
	if err := db.DB.Where("from_tx_hash =  ?", txHash).First(&rcd).Error; err != nil {
		return nil, err
	}
	return &rcd, nil
}

// Withdraw
func (db *Database) CreateWithdrawTransaction(tx *spec.ZuluWithdrawInfo) error {
	return db.DB.Create(tx).Error
}

func (db *Database) CreateAndGetWithdrawTransaction(tx *spec.ZuluWithdrawInfo) (*spec.ZuluWithdrawInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.ZuluWithdrawInfo
	if err := db.DB.Where("from_tx_hash = ?", tx.FromTxHash).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) ListWithdrawTransactionByAddress(toAddress string) (*[]spec.ZuluWithdrawInfo, error) {
	var txInfos []spec.ZuluWithdrawInfo
	if err := db.DB.Where("to_address = ?", toAddress).Order("created_at DESC").Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) ListUnhandledWithdrawTransactions(number int) (*[]spec.ZuluWithdrawInfo, error) {
	var txInfos []spec.ZuluWithdrawInfo
	if err := db.DB.Where("state = ?", spec.WithdrawTxStatePending).Order("created_at asc").Limit(number).Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) UpdateWithdrawTransaction(tx *spec.ZuluWithdrawInfo) error {
	return db.DB.Save(tx).Error
}

///Deprecated
func (db *Database) GetDepositTransactionOld(txid, fromAddress string) (*spec.ZuluDepositTxInfo, error) {
	var txInfo spec.ZuluDepositTxInfo
	if err := db.DB.Where("transaction_id = ? AND l1_address = ?", txid, fromAddress).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) CreateDepositTransactionOld(tx *spec.ZuluDepositTxInfo) error {
	return db.DB.Create(tx).Error
}

func (db *Database) CreateAndGetDepositTransactionOld(tx *spec.ZuluDepositTxInfo) (*spec.ZuluDepositTxInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.ZuluDepositTxInfo
	if err := db.DB.Where("transaction_id = ? AND l1_address = ?", tx.TransactionID, tx.L1Address).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) ListDepositTransactionByAddressOld(fromAddress string) (*[]spec.ZuluDepositTxInfo, error) {
	var txInfos []spec.ZuluDepositTxInfo
	if err := db.DB.Where("l1_address = ?", fromAddress).Order("created_at DESC").Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) ListUnhandledDepositTransactionsOld(number int) (*[]spec.ZuluDepositTxInfo, error) {
	var depositTxs []spec.ZuluDepositTxInfo
	if err := db.DB.Where("state = ?", spec.DepositTxStateProcessing).Order("created_at asc").Limit(number).Find(&depositTxs).Error; err != nil {
		return nil, err
	}
	return &depositTxs, nil
}

func (db *Database) UpdateDepositTransactionOld(tx *spec.ZuluDepositTxInfo) error {
	return db.DB.Save(tx).Error
}

func (db *Database) DeleteDepositTransactionOld(txid, fromAddress string) error {
	var txInfo spec.ZuluDepositTxInfo
	if err := db.DB.Where("transaction_id = ? AND l1_address = ?", txid, fromAddress).Delete(&txInfo).Error; err != nil {
		return err
	}
	return nil
}

// withdraw
func (db *Database) GetZuluWithdrawTransactionOld(txhash, fromAddress string) (*spec.ZuluWithdrawTxInfo, error) {
	var txInfo spec.ZuluWithdrawTxInfo
	if err := db.DB.Where("tx_hash = ? AND l2_address = ?", txhash, fromAddress).Find(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) CreateZuluWithdrawTransactionOld(tx *spec.ZuluWithdrawTxInfo) error {
	return db.DB.Create(tx).Error
}

func (db *Database) CreateAndGetZuluWithdrawTransactionOld(tx *spec.ZuluWithdrawTxInfo) (*spec.ZuluWithdrawTxInfo, error) {
	if err := db.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	var txInfo spec.ZuluWithdrawTxInfo
	if err := db.DB.Where("tx_hash = ? AND l2_address = ?", tx.TxHash, tx.L2Address).First(&txInfo).Error; err != nil {
		return nil, err
	}
	return &txInfo, nil
}

func (db *Database) ListZuluWithdrawTransactionByAddressOld(fromAddress string) (*[]spec.ZuluWithdrawTxInfo, error) {
	var txInfos []spec.ZuluWithdrawTxInfo
	if err := db.DB.Where("l2_address = ?", fromAddress).Order("created_at DESC").Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) ListUnhandledWithdrawTransactionsOld(number int) (*[]spec.ZuluWithdrawTxInfo, error) {
	var txInfos []spec.ZuluWithdrawTxInfo
	if err := db.DB.Where("state = ?", spec.WithdrawTxStatePending).Order("created_at asc").Limit(number).Find(&txInfos).Error; err != nil {
		return nil, err
	}
	return &txInfos, nil
}

func (db *Database) UpdateZuluWithdrawTransactionOld(tx *spec.ZuluWithdrawTxInfo) error {
	return db.DB.Save(tx).Error
}

func (db *Database) DeleteZuluWithdrawTransactionOld(txhash, fromAddress string) error {
	var txInfo spec.ZuluWithdrawTxInfo
	if err := db.DB.Where("tx_hash = ? AND l2_address = ?", txhash, fromAddress).Delete(&txInfo).Error; err != nil {
		return err
	}
	return nil
}
