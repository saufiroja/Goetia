package mappers

import (
	"encoding/json"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"net/http"
)

func NewResponseMapper(w http.ResponseWriter, statusCode int, message string, response any) error {
	res := responses.NewResponse(message, statusCode, response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return responses.NewResponse(err.Error(), http.StatusInternalServerError, response)
	}

	return nil
}
