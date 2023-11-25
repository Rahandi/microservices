package services

import (
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/models"
	"context"
	"fmt"
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

func (s *FinancialService) AccountCreate(ctx context.Context, input *models.AccountCreateInput) error {
	IAMData, err := s.redis.GetUserIAMData(ctx)
	if err != nil {
		return err
	}

	payload := &models.FinancialServiceAccountCreateRequest{
		UserId:        IAMData.Id,
		Name:          input.Name,
		AccountNumber: input.AccountNumber,
	}

	fmt.Println(payload)

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
