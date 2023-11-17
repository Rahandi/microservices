package services

import (
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/types"
	"context"
)

func GetAuthHeader(ctx context.Context, redis *internals.Redis) (map[string]string, error) {
	accountId := ctx.Value(types.AccountIdKey).(string)
	token, err := redis.Client.HGet(ctx, accountId, "token").Result()
	if err != nil {
		return nil, err
	}

	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}

	return headers, nil
}
