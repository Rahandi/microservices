package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type WhoAmIOutput struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Principal string    `json:"principal"`
}

type JWTClaims struct {
	*jwt.StandardClaims
	Principal string `json:"principal"`
}

type JWTRefreshClaims struct {
	*jwt.StandardClaims
	Token string `json:"token"`
}
