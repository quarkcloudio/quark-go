package dal

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/dal/db"
	redisclient "github.com/quarkcloudio/quark-go/v2/pkg/dal/redis"
	"github.com/redis/go-redis/v9"
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
