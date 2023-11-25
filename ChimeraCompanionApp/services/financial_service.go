package services

import (
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/models"
	"ChimeraCompanionApp/types"
	"context"

	"github.com/google/uuid"
)

type FinancialService struct {
	http  *internals.Http
	redis *internals.Redis
}

func NewFinancialService(config *internals.Config, redis *internals.Redis) *FinancialService {
	httpClient := internals.NewHttp(config.FinancialServiceEndpoint)

	return &FinancialService{
		http:  httpClient,
		redis: redis,
	}
}

func (s *FinancialService) CreateAccount(ctx context.Context, input *models.AccountCreateInput) error {
	accountId := ctx.Value(types.AccountIdKey).(string)
	userId, err := s.redis.Client.HGet(ctx, accountId, "iamservice.id").Result()
	if err != nil {
		return err
	}

	payload := &models.FinancialServiceAccountCreateRequest{
		UserId:        uuid.MustParse(userId),
		Name:          input.Name,
		AccountNumber: input.AccountNumber,
	}

	response := &models.FinancialServiceAccountCreateResponse{}
	err = s.http.Post(ctx, "/account/create", payload, response, nil)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return err
	}

	return nil
}
