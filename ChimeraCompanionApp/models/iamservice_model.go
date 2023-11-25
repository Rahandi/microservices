package models

import "github.com/google/uuid"

type RegisterInput struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AccountId string `json:"account_id"`
	Password  string `json:"password"`
}

type IAMServiceRegisterRequest struct {
	Name      string `json:"name"`
	Principal string `json:"principal"`
	Password  string `json:"password"`
}

type IAMServiceRegisterResponse struct {
	Data struct {
		Id           uuid.UUID `json:"id"`
		Token        string    `json:"token"`
		RefreshToken string    `json:"refresh_token"`
	}
	Error string `json:"error"`
}

type LoginInput struct {
	AccountId string `json:"account_id"`
	Password  string `json:"password"`
}

type IAMServiceLoginRequest struct {
	Principal string `json:"principal"`
	Password  string `json:"password"`
}

type IAMServiceLoginResponse struct {
	Data struct {
		Id           uuid.UUID `json:"id"`
		Token        string    `json:"token"`
		RefreshToken string    `json:"refresh_token"`
	}
	Error string `json:"error"`
}

type IAMServiceWhoAmIResponse struct {
	Data struct {
		Id        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Principal string    `json:"principal"`
	}
	Error string `json:"error"`
}

type IAMServiceRefreshTokenRequest struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type IAMServiceRefreshTokenResponse struct {
	Data struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	Error string `json:"error"`
}
