// handler_utils.go
package handlers

import (
	"encoding/json"
	"iamService/models"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	errorResponse := &models.ErrorResponse{
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(errorResponse)
}
