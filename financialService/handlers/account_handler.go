package handlers

import (
	"FinancialService/models"
	"FinancialService/services"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) Register(httpServer *http.ServeMux) {
	httpServer.HandleFunc("/accounts", h.ResourceHandler)
	httpServer.HandleFunc("/accounts/", h.SingleResourceHandler)
}

func (h *AccountHandler) ResourceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.List(w, r)
	case http.MethodPost:
		h.Create(w, r)
	case http.MethodPut:
		h.Update(w, r)
	default:
		handleError(w, errors.New("method not allowed"))
	}
}

func (h *AccountHandler) SingleResourceHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/account/")
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r, id)
	case http.MethodDelete:
		h.Delete(w, r, id)
	default:
		handleError(w, errors.New("method not allowed"))
	}
}

func (h *AccountHandler) List(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accountService.List()
	if err != nil {
		handleError(w, err)
		return
	}
	handleSuccess(w, accounts)
}

func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request, id string) {
	account, err := h.accountService.Get(id)
	if err != nil {
		handleError(w, err)
		return
	}
	handleSuccess(w, account)
}

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

func (h *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.AccountUpdateInput
	err := decoder.Decode(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	err = h.accountService.Update(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	handleSuccess(w, "Account updated successfully")
}

func (h *AccountHandler) Delete(w http.ResponseWriter, r *http.Request, id string) {
	err := h.accountService.Delete(id)
	if err != nil {
		handleError(w, err)
		return
	}
	handleSuccess(w, "Account deleted successfully")
}
