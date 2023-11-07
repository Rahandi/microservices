package models

import "github.com/golang-jwt/jwt"

type WhoAmIResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type JWTClaims struct {
	*jwt.StandardClaims
	Email string `json:"email"`
}
