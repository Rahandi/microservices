package models

type RegisterInput struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	AccountId string `json:"account_id"`
}

type IAMServiceRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IAMServiceRegisterResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
