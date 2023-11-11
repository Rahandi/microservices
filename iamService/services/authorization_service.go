package services

import (
	"errors"
	"IAMlService/internals"
	"IAMlService/models"
	"IAMlService/repositories"

	"github.com/google/uuid"
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

func (s *AuthorizationService) WhoAmI(token string) (*models.WhoAmIOutput, error) {
	claims, err := s.authenticationService.DecodeToken(token)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, err
	}

	user := s.userRepository.FindByID(userId)
	if user == nil {
		return nil, errors.New("user not found")
	}

	return &models.WhoAmIOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
