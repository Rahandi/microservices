package models

import "github.com/google/uuid"

type RegisterInput struct {
	Name      string `json:"name"`
	Principal string `json:"principal"`
	Password  string `json:"password"`
}

type RegisterOutput struct {
	Id           uuid.UUID `json:"id"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

type LoginInput struct {
	Principal string `json:"principal"`
	Password  string `json:"password"`
}

type LoginOutput struct {
	Id           uuid.UUID `json:"id"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

type RefreshTokenInput struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenOutput struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
