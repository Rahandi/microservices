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
	Data struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	Error string `json:"error"`
}

type LoginInput struct {
	Email     string `json:"email"`
	AccountId string `json:"account_id"`
}

type IAMServiceLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IAMServiceLoginResponse struct {
	Data struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	Error string `json:"error"`
}
