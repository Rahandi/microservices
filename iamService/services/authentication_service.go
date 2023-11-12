package services

import (
	"IAMService/internals"
	"IAMService/models"
	"IAMService/repositories"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

func (s *AuthenticationService) Register(input *models.RegisterInput) (*models.RegisterOutput, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:      input.Name,
		Principal: input.Principal,
		Password:  string(hashedPassword),
	}

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(user, token)
	if err != nil {
		return nil, err
	}

	return &models.RegisterOutput{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthenticationService) Login(input *models.LoginInput) (*models.LoginOutput, error) {
	user := s.userRepository.FindByPrincipal(input.Principal)
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

	refreshToken, err := s.generateRefreshToken(user, token)
	if err != nil {
		return nil, err
	}

	return &models.LoginOutput{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthenticationService) RefreshToken(input *models.RefreshTokenInput) (*models.RefreshTokenOutput, error) {
	claims, err := s.DecodeToken(input.Token)
	if err != nil {
		return nil, err
	}

	refreshClaims, err := s.DecodeRefreshToken(input.RefreshToken)
	if err != nil {
		return nil, err
	}

	if claims.Subject != refreshClaims.Subject || input.Token != refreshClaims.Token {
		return nil, errors.New("invalid token")
	}

	userId, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, err
	}

	user := s.userRepository.FindByID(userId)
	if user == nil {
		return nil, errors.New("user not found")
	}

	newToken, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.generateRefreshToken(user, newToken)
	if err != nil {
		return nil, err
	}

	return &models.RefreshTokenOutput{
		Token:        newToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *AuthenticationService) generateToken(user *models.User) (string, error) {
	parsedExpires, err := time.ParseDuration(s.config.JwtExpires)
	if err != nil {
		log.Panic(err)
	}

	claims := &models.JWTClaims{
		StandardClaims: &jwt.StandardClaims{
			Subject:   fmt.Sprint(user.ID),
			Issuer:    "IAMService",
			ExpiresAt: time.Now().Add(parsedExpires).Unix(),
		},
		Principal: user.Principal,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.config.JwtSecret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (s *AuthenticationService) generateRefreshToken(user *models.User, token string) (string, error) {
	parsedExpires, err := time.ParseDuration(s.config.JwtRefreshExpires)
	if err != nil {
		log.Panic(err)
	}

	claims := &models.JWTRefreshClaims{
		StandardClaims: &jwt.StandardClaims{
			Subject:   fmt.Sprint(user.ID),
			Issuer:    "IAMService",
			ExpiresAt: time.Now().Add(parsedExpires).Unix(),
		},
		Token: token,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := refreshToken.SignedString([]byte(s.config.JwtRefreshSecret))
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

func (s *AuthenticationService) DecodeRefreshToken(token string) (*models.JWTRefreshClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &models.JWTRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JwtRefreshSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(*models.JWTRefreshClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
