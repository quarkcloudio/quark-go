package db

import (
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
}
