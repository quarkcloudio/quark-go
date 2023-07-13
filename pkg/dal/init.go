package dal

import (
	"github.com/go-redis/redis/v8"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	redisclient "github.com/quarkcms/quark-go/v2/pkg/dal/redis"
	"gorm.io/gorm"
)

// Init init DB
func InitDB(dialector gorm.Dialector, opts gorm.Option) {
	db.Init(dialector, opts)
}

// Init init redis
func InitRedis(options *redis.Options) {
	redisclient.Init(options)
}
