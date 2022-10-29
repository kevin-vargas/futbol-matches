package redis

import (
	"github.com/go-redis/redis/v8"
)

func GetConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "futbol-matches-metrics:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
