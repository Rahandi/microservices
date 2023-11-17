package services

import "ChimeraCompanionApp/internals"

type FinancialService struct {
	http *internals.Http
	redis *internals.Redis
}

func NewFinancialService(config *internals.Config, redis *internals.Redis) *FinancialService {
	httpClient := internals.NewHttp(config.FinancialServiceEndpoint)

	return &FinancialService{
		http: httpClient,
		redis: redis,
	}
}