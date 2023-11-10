package models

import "github.com/google/uuid"

type AccountCreateInput struct {
	UserId        uuid.UUID `json:"user_id"`
	Name          string    `json:"name"`
	AccountNumber string    `json:"account_number"`
}
