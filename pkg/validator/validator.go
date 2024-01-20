package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Validation struct {
	validation validator.Validate
}

func NewValidation() *Validation {
	return &Validation{
		validation: *validator.New(),
	}
}

func (v *Validation) Validate(i interface{}) error {
	return v.validation.Struct(i)
}

func (v *Validation) CustomError(err error) string {
	var errMessage string
	for _, v := range err.(validator.ValidationErrors) {
		switch v.Tag() {
		case "required":
			errMessage = fmt.Sprintf("%s is required", strings.ToLower(v.Field()))
		}
	}

	return errMessage
}
