package service

import (
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/models"
)

type IAMService struct {
	httpClient *internals.HttpClient
	config     *internals.Config
}

func NewIAMService(config *internals.Config) *IAMService {
	httpClient := internals.NewHttpClient(config.IAMServiceEndpoint)

	return &IAMService{
		httpClient: httpClient,
		config:     config,
	}
}

func (s *IAMService) Register(input *models.RegisterInput) error {
	payload := &models.IAMServiceRegisterRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: "password",
	}

	response := &models.IAMServiceRegisterResponse{}

	err := s.httpClient.Post("/register", payload, response)
	if err != nil {
		return err
	}

	return nil
}
