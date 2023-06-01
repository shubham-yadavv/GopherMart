package database

import (
	redis "github.com/go-redis/redis/v8"

	"github.com/shubham-yadavv/go-ecommerce/config"
)

var redisClient *redis.Client

func ConnectRedis() *redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{

			Addr: config.GetEnv("REDIS_HOST") + ":" + config.GetEnv("REDIS_PORT"),
			// Password: config.GetEnv("REDIS_PASSWORD"),
			DB: 0,
		})
	}
	return redisClient
}

func GetRedisClient() *redis.Client {
	return redisClient
}
