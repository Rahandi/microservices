package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id" redis:"id"`
	Username string `json:"username" redis:"username"`

	IAMData IAMData `json:"iamdata" redis:"iamdata"`
}

type IAMData struct {
	Id           uuid.UUID `json:"id" redis:"id"`
	Token        string    `json:"token" redis:"token"`
	RefreshToken string    `json:"refresh_token" redis:"refresh_token"`
}

func (m IAMData) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
