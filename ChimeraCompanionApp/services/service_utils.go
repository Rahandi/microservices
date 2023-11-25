package services

import (
	"ChimeraCompanionApp/internals"
	"context"
)

func GetAuthHeader(ctx context.Context, redis *internals.Redis) (map[string]string, error) {
	IAMData, err := redis.GetUserIAMData(ctx)
	if err != nil {
		return nil, err
	}
	token := IAMData.Token

	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}

	return headers, nil
}
