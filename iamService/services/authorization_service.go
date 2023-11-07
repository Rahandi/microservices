package services

import (
	"errors"
	"iamService/internals"
	"iamService/models"
	"iamService/repositories"
	"strconv"
)

type AuthorizationService struct {
	config                *internals.Config
	userRepository        *repositories.UserRepository
	authenticationService *AuthenticationService
}

func NewAuthorizationService(config *internals.Config, repository *repositories.UserRepository, authenticationService *AuthenticationService) *AuthorizationService {
	return &AuthorizationService{
		config:                config,
		userRepository:        repository,
		authenticationService: authenticationService,
	}
}

func (s *AuthorizationService) WhoAmI(token string) (*models.WhoAmIResponse, error) {
	claims, err := s.authenticationService.DecodeToken(token)
	if err != nil {
		return nil, err
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return nil, err
	}

	user := s.userRepository.FindByID(uint(userId))
	if user == nil {
		return nil, errors.New("user not found")
	}

	return &models.WhoAmIResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
