package handlers

import (
	"encoding/json"
	"iamService/models"
	"iamService/services"
	"net/http"
)

type AuthenticationHandler struct {
	httpServer  *http.ServeMux
	authService *services.AuthenticationService
}

func NewAuthenticationHandler(httpServer *http.ServeMux, authService *services.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		httpServer:  httpServer,
		authService: authService,
	}
}

func (h *AuthenticationHandler) Register() {
	h.httpServer.HandleFunc("/register", h.registerHandler)
	h.httpServer.HandleFunc("/login", h.loginHandler)
	h.httpServer.HandleFunc("/whoami", h.whoamiHandler)
}

func (h *AuthenticationHandler) handleError(err error, w http.ResponseWriter) {
	errorResponse := &models.ErrorResponse{
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(errorResponse)
}

func (h *AuthenticationHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.RegisterRequest
	err := decoder.Decode(&input)
	if err != nil {
		h.handleError(err, w)
		return
	}
	response, err := h.authService.Register(&input)
	if err != nil {
		h.handleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (h *AuthenticationHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.LoginRequest
	err := decoder.Decode(&input)
	if err != nil {
		h.handleError(err, w)
		return
	}
	response, err := h.authService.Login(&input)
	if err != nil {
		h.handleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (h *AuthenticationHandler) whoamiHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	token = token[7:]
	response, err := h.authService.WhoAmI(token)
	if err != nil {
		h.handleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(response)
}
