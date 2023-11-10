// handler_utils.go
package handlers

import (
	"encoding/json"
	"financialService/models"
	"net/http"
)

func handleSuccess(w http.ResponseWriter, response interface{}) {
	successResponse := &models.SuccessResponse{
		Data: response,
	}
	json.NewEncoder(w).Encode(successResponse)
}

func handleError(w http.ResponseWriter, err error) {
	errorResponse := &models.ErrorResponse{
		Error: err.Error(),
	}
	json.NewEncoder(w).Encode(errorResponse)
}
