package internals

import (
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(config *Config) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		DB:       config.RedisDB,
		Password: config.RedisPassword,
	})

	return RedisClient{
		client: client,
	}
}
