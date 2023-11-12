package services

import (
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/models"
	"errors"
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

func (s *IAMService) Register(input *models.RegisterInput) (*models.IAMServiceRegisterResponse, error) {
	payload := &models.IAMServiceRegisterRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.AccountId,
	}

	response := &models.IAMServiceRegisterResponse{}

	err := s.httpClient.Post("/register", payload, response)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	return response, nil
}

func (s *IAMService) Login(input *models.LoginInput) (*models.IAMServiceLoginResponse, error) {
	payload := &models.IAMServiceLoginRequest{
		Email:    input.Email,
		Password: input.AccountId,
	}

	response := &models.IAMServiceLoginResponse{}

	err := s.httpClient.Post("/login", payload, response)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	return response, nil
}
