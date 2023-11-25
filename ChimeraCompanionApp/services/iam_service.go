package services

import (
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/models"
	"ChimeraCompanionApp/types"
	"context"
	"errors"
	"fmt"
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

func (s *IAMService) Register(ctx context.Context, input *models.RegisterInput) error {
	payload := &models.IAMServiceRegisterRequest{
		Name:      input.Name,
		Principal: input.AccountId,
		Password:  input.Password,
	}

	response := &models.IAMServiceRegisterResponse{}
	err := s.http.Post(ctx, "/register", payload, response, nil)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return errors.New(response.Error)
	}

	accountId := input.AccountId
	user := models.User{
		ID:       accountId,
		Username: input.Username,
		IAMData: models.IAMData{
			Id:           response.Data.Id,
			Token:        response.Data.Token,
			RefreshToken: response.Data.RefreshToken,
		},
	}
	_, err = s.redis.Client.HSet(ctx, accountId, user).Result()
	if err != nil {
		return err
	}

	return nil
}

func (s *IAMService) Login(ctx context.Context, input *models.LoginInput) error {
	payload := &models.IAMServiceLoginRequest{
		Principal: input.AccountId,
		Password:  input.Password,
	}

	response := &models.IAMServiceLoginResponse{}
	err := s.http.Post(ctx, "/login", payload, response, nil)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return errors.New(response.Error)
	}

	fmt.Println(response.Data.Id)

	accountId := input.AccountId
	user := models.User{
		ID:       accountId,
		Username: input.Password,
		IAMData: models.IAMData{
			Id:           response.Data.Id,
			Token:        response.Data.Token,
			RefreshToken: response.Data.RefreshToken,
		},
	}
	_, err = s.redis.Client.HSet(ctx, accountId, user).Result()
	if err != nil {
		return err
	}

	return nil
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

func (s *IAMService) RefreshToken(ctx context.Context) error {
	accountId := ctx.Value(types.AccountIdKey).(string)
	token, err := s.redis.Client.HGet(ctx, accountId, "token").Result()
	if err != nil {
		return err
	}

	refreshToken, err := s.redis.Client.HGet(ctx, accountId, "refresh_token").Result()
	if err != nil {
		return err
	}

	request := &models.IAMServiceRefreshTokenRequest{
		Token:        token,
		RefreshToken: refreshToken,
	}
	response := &models.IAMServiceRefreshTokenResponse{}
	err = s.http.Post(ctx, "/refresh-token", request, response, nil)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return errors.New(response.Error)
	}

	_, err = s.redis.Client.HSet(ctx, accountId, "token", response.Data.Token, "refresh_token", response.Data.RefreshToken).Result()
	if err != nil {
		return err
	}

	return nil
}
