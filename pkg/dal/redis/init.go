package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func Init(options *redis.Options) {
	Client = redis.NewClient(options)
	if err := Client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
