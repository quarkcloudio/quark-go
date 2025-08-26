package login

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CaptchaStore struct {
	RedisClient *redis.Client
	Expiration  time.Duration
}

func (store *CaptchaStore) Set(id string, digits []byte) {
	store.RedisClient.Set(context.Background(), id, string(digits), store.Expiration)
}

func (store *CaptchaStore) Get(id string, clear bool) (digits []byte) {
	bytes, _ := store.RedisClient.Get(context.Background(), id).Bytes()
	return bytes
}
