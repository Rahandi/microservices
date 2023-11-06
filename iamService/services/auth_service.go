package services

import (
	"errors"
	"iamService/internals"
	"iamService/models"
	"iamService/repositories"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	config     *internals.Config
	repository *repositories.AuthRepository
}

func NewAuthService(config *internals.Config, repository *repositories.AuthRepository) *AuthService {
	return &AuthService{
		config:     config,
		repository: repository,
	}
}

func (s *AuthService) Register(input *models.RegisterRequest) (*models.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	err = s.repository.Create(&user)
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

func (s *AuthService) Login(input *models.LoginRequest) (*models.LoginResponse, error) {
	user := s.repository.FindByEmail(input.Email)
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

func (s *AuthService) WhoAmI(token string) (*models.WhoAmIResponse, error) {
	claims, err := s.decodeToken(token)
	if err != nil {
		return nil, err
	}

	sub := claims["sub"].(float64)
	user := s.repository.FindByID(uint(sub))
	if user == nil {
		return nil, errors.New("user not found")
	}

	return &models.WhoAmIResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *AuthService) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.config.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (s *AuthService) decodeToken(token string) (jwt.MapClaims, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
