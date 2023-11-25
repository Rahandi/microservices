package models

import "github.com/google/uuid"

type AccountCreateInput struct {
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
}

type FinancialServiceAccountCreateRequest struct {
	UserId        uuid.UUID `json:"user_id"`
	Name          string    `json:"name"`
	AccountNumber string    `json:"account_number"`
}

type FinancialServiceAccountCreateResponse struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}
