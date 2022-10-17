package config

import (
	"context"

	"github.com/go-redis/redis/v9"
)

var Rdb *redis.Client

var Ctx = context.TODO()

func ConnectRedis() (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     ConfigApp.Redis.Address,
		Password: ConfigApp.Redis.Password,
		DB:       0, // use default DB
	})

	if _, err = rdb.Ping(Ctx).Result(); err != nil {
		return
	}

	return
}
