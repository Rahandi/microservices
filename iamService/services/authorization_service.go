package services

import (
	"iamService/internals"
	"iamService/models"
	"iamService/repositories"
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

	sub := claims["sub"].(float64)
	user := s.userRepository.FindByID(uint(sub))
	if user == nil {
		return nil, nil
	}

	return &models.WhoAmIResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
