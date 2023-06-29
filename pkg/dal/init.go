package dal

import (
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/dal/redis"
	"gorm.io/gorm"
)

// Init init DB
func InitDB(dialector gorm.Dialector, opts gorm.Option) {
	db.Init(dialector, opts)
}

// Init init redis
func InitRedis(addr string, password string) {
	redis.Init(addr, password)
}
