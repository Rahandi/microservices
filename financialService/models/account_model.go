package models

import "github.com/google/uuid"

type AccountCreateInput struct {
	UserId  uuid.UUID `json:"user_id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
}

type AccountUpdateInput struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
}
