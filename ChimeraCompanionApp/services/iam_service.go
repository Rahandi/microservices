package services

import (
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/models"
	"context"
	"errors"
)

type IAMService struct {
	http  *internals.Http
	redis *internals.Redis
}

func NewIAMService(config *internals.Config, redis *internals.Redis) *IAMService {
	httpClient := internals.NewHttp(config.IAMServiceEndpoint)

	return &IAMService{
		http:  httpClient,
		redis: redis,
	}
}

func (s *IAMService) Register(ctx context.Context, input *models.RegisterInput) (*models.IAMServiceRegisterResponse, error) {
	payload := &models.IAMServiceRegisterRequest{
		Name:      input.Name,
		Principal: input.AccountId,
		Password:  input.Password,
	}

	response := &models.IAMServiceRegisterResponse{}
	err := s.http.Post(ctx, "/register", payload, response, nil)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	accountId := input.AccountId
	user := models.User{
		ID:           accountId,
		Username:     input.Username,
		Token:        response.Data.Token,
		RefreshToken: response.Data.RefreshToken,
	}
	_, err = s.redis.Client.HSet(ctx, accountId, user).Result()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *IAMService) Login(ctx context.Context, input *models.LoginInput) (*models.IAMServiceLoginResponse, error) {
	payload := &models.IAMServiceLoginRequest{
		Principal: input.AccountId,
		Password:  input.Password,
	}

	response := &models.IAMServiceLoginResponse{}
	err := s.http.Post(ctx, "/login", payload, response, nil)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	accountId := input.AccountId
	exists, err := s.redis.Client.HExists(ctx, accountId, "token").Result()
	if err != nil {
		return nil, err
	}
	if !exists {
		user := models.User{
			ID:           accountId,
			Username:     input.Password,
			Token:        response.Data.Token,
			RefreshToken: response.Data.RefreshToken,
		}
		_, err = s.redis.Client.HSet(ctx, accountId, user).Result()
		if err != nil {
			return nil, err
		}
	} else {
		_, err = s.redis.Client.HSet(ctx, accountId, "token", response.Data.Token, "refresh_token", response.Data.RefreshToken).Result()
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (s *IAMService) WhoAmI(ctx context.Context) (*models.IAMServiceWhoAmIResponse, error) {
	headers, err := GetAuthHeader(ctx, s.redis)
	if err != nil {
		return nil, err
	}

	response := &models.IAMServiceWhoAmIResponse{}
	err = s.http.Get(ctx, "/whoami", response, &headers)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	return response, nil
}
