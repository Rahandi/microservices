package internals

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(config *Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		DB:       config.RedisDB,
		Password: config.RedisPassword,
	})

	return &Redis{
		Client: client,
	}
}
