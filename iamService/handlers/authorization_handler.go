package handlers

import (
	"iamService/services"
	"net/http"
)

type AuthorizationHandler struct {
	authorizationService *services.AuthorizationService
}

func NewAuthorizationHandler(authorizationService *services.AuthorizationService) *AuthorizationHandler {
	return &AuthorizationHandler{
		authorizationService: authorizationService,
	}
}

func (h *AuthorizationHandler) Register(httpServer *http.ServeMux) {
	httpServer.HandleFunc("/whoami", h.whoamiHandler)
}

func (h *AuthorizationHandler) whoamiHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")[7:]
	response, err := h.authorizationService.WhoAmI(token)
	if err != nil {
		handleError(w, err)
		return
	}
	handleSuccess(w, response)
}
