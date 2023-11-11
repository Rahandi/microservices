package handlers

import (
	"IAMlService/models"
	"IAMlService/services"
	"encoding/json"
	"net/http"
)

type AuthenticationHandler struct {
	authenticationService *services.AuthenticationService
}

func NewAuthenticationHandler(authenticationService *services.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		authenticationService: authenticationService,
	}
}

func (h *AuthenticationHandler) Register(httpServer *http.ServeMux) {
	httpServer.HandleFunc("/register", h.registerHandler)
	httpServer.HandleFunc("/login", h.loginHandler)
	httpServer.HandleFunc("/refresh-token", h.refreshTokenHandler)
}

func (h *AuthenticationHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.RegisterInput
	err := decoder.Decode(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	response, err := h.authenticationService.Register(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	handleSuccess(w, response)
}

func (h *AuthenticationHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.LoginInput
	err := decoder.Decode(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	response, err := h.authenticationService.Login(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	handleSuccess(w, response)
}

func (h *AuthenticationHandler) refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.RefreshTokenInput
	err := decoder.Decode(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	response, err := h.authenticationService.RefreshToken(&input)
	if err != nil {
		handleError(w, err)
		return
	}
	handleSuccess(w, response)
}
