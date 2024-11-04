package config

import (
	"exchangeapp/global"
	"log"

	"github.com/go-redis/redis"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "1234",
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error")
	}

	global.RedisDB = RedisClient
}
