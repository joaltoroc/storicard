package dtos

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
)

type Payload struct {
	FileName string `json:"fileName"`
	Email    string `json:"email"`
}

func (pay Payload) Validate() error {
	return validation.ValidateStruct(&pay,
		validation.Field(&pay.FileName, validation.Required),
		validation.Field(&pay.Email, validation.Required, is.Email),
	)
}
