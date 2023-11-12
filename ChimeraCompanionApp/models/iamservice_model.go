package models

type RegisterInput struct {
	Name      string `json:"name"`
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
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
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
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	Error string `json:"error"`
}
