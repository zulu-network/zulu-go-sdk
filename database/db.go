package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/zulu-network/zulu-go-sdk/log"
)

type Database struct {
	DB  *gorm.DB
	log *log.Logger
}

func NewDB(url string) (*Database, error) {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{
		DB:  db,
		log: log.L().With(log.Any("service", "db")),
	}, nil
}
