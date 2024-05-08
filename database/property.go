package database

import (
	"github.com/zulu-network/zulu-go-sdk/spec"
)

func (db *Database) GetProperty(name string) (string, error) {
	var prop spec.ZuluProperty
	if err := db.DB.Where("name = ?", name).First(&prop).Error; err != nil {
		return "", err
	}
	return prop.Value, nil
}

func (db *Database) UpdateProperty(prop *spec.ZuluProperty) error {
	return db.DB.Model(&spec.ZuluProperty{}).Where("name = ?", prop.Name).
		Update("value", prop.Value).Error
}

func (db *Database) UpdateAndGetProperty(prop *spec.ZuluProperty) (*spec.ZuluProperty, error) {
	tx := db.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&spec.ZuluProperty{}).Where("name = ?", prop.Name).
		Update("value", prop.Value).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var newProp spec.ZuluProperty
	if err := tx.Where("name = ?", prop.Name).First(&newProp).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &newProp, nil
}
