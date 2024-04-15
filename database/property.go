package database

import (
	"github.com/zulu-network/zulu-go-sdk/spec"
)

func (db *Database) GetProperty(name string) (string, error) {
	var prop spec.ZuluProperty
	if err := db.DB.Where("name = ?", name).Find(&prop).Error; err != nil {
		return "", err
	}
	return prop.Value, nil
}

func (db *Database) UpdateProperty(prop *spec.ZuluProperty) error {
	if err := db.DB.Model(&spec.ZuluProperty{}).Where("name = ?", prop.Name).
		Update("value", prop.Value).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateAndGetProperty(prop *spec.ZuluProperty) (*spec.ZuluProperty, error) {
	if err := db.DB.Model(&spec.ZuluProperty{}).Where("name = ?", prop.Name).
		Update("value", prop.Value).Error; err != nil {
		return nil, err
	}
	var newProp spec.ZuluProperty
	if err := db.DB.Where("name = ?", prop.Name).Find(&newProp).Error; err != nil {
		return nil, err
	}
	return &newProp, nil
}
