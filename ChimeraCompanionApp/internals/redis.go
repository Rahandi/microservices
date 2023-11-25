package internals

import (
	"ChimeraCompanionApp/models"
	"ChimeraCompanionApp/types"
	"context"
	"encoding/json"

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

func (r *Redis) GetUserIAMData(ctx context.Context) (*models.IAMData, error) {
	accountId := ctx.Value(types.AccountIdKey).(string)
	data, err := r.Client.HGet(ctx, accountId, "iamdata").Result()
	if err != nil {
		return nil, err
	}

	IAMData := &models.IAMData{}
	err = json.Unmarshal([]byte(data), IAMData)
	if err != nil {
		return nil, err
	}

	return IAMData, nil
}
