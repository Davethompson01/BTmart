package Validation

import (
	"bsmart/internals/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateAccount(auth models.MartAuth) error {
	if err := validate.Struct(auth); err != nil {
		return FormatValidationError(err)
	}
	return nil
}

