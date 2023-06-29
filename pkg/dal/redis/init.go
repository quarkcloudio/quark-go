package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func Init(addr string, password string) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})
	if err := Client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
