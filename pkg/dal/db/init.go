package db

import (
	"time"

	"gorm.io/gorm"
)

var Client *gorm.DB

// Init init DB
func Init(dialector gorm.Dialector, opts gorm.Option) {
	var err error
	Client, err = gorm.Open(dialector, opts)
	if err != nil {
		panic(err)
	}

	sqlDB, err := Client.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Minute * 2)
}
