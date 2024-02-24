package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Init(options *redis.Options) {
	Client = redis.NewClient(options)
	if err := Client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
