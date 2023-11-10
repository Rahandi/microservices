package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type WhoAmIOutput struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type JWTClaims struct {
	*jwt.StandardClaims
	Email string `json:"email"`
}

type JWTRefreshClaims struct {
	*jwt.StandardClaims
	Token string `json:"token"`
}
