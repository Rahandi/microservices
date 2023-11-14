package models

type User struct {
	ID       string `json:"id" redis:"id"`
	Username string `json:"username" redis:"username"`

	// IAM Service
	Token        string `json:"token" redis:"token"`
	RefreshToken string `json:"refresh_token" redis:"refresh_token"`
}
