package handlers

import (
	"encoding/json"
	"financialService/models"
	"financialService/services"
	"net/http"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) Register(httpServer *http.ServeMux) {}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.AccountCreateInput
	err := decoder.Decode(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	err = h.accountService.Create(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	handleSuccess(w, "Account created successfully")
}