package services

import (
	"errors"
	"fmt"
	"iamService/internals"
	"iamService/models"
	"iamService/repositories"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService struct {
	config         *internals.Config
	userRepository *repositories.UserRepository
}

func NewAuthenticationService(config *internals.Config, repository *repositories.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		config:         config,
		userRepository: repository,
	}
}

func (s *AuthenticationService) Register(input *models.RegisterRequest) (*models.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}

	user := models.DBUser{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	err = s.userRepository.Create(&user)
	if err != nil {
		return nil, err
	}

	token, err := s.generateToken(&user)
	if err != nil {
		return nil, err
	}

	return &models.RegisterResponse{
		Token: token,
	}, nil
}

func (s *AuthenticationService) Login(input *models.LoginRequest) (*models.LoginResponse, error) {
	user := s.userRepository.FindByEmail(input.Email)
	if user == nil {
		return nil, errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, err
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
	}, nil
}

func (s *AuthenticationService) generateToken(user *models.DBUser) (string, error) {
	parsedExpires, err := time.ParseDuration(s.config.JwtExpires)
	if err != nil {
		log.Panic(err)
	}

	claims := &models.JWTClaims{
		StandardClaims: &jwt.StandardClaims{
			Subject:   fmt.Sprint(user.ID),
			Issuer:    "iamService",
			ExpiresAt: time.Now().Add(parsedExpires).Unix(),
		},
		Email: user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.config.JwtSecret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (s *AuthenticationService) DecodeToken(token string) (*models.JWTClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(*models.JWTClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
