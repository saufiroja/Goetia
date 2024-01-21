package mappers

import (
	"github.com/saufiroja/cqrs/internal/contracts/responses"
)

func NewResponseMapper(statusCode int, message string, response any) error {
	return responses.NewResponse(message, statusCode, response)
}
