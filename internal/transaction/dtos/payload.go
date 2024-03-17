package dtos

import "github.com/invopop/validation"

type Payload struct {
	FileName string `json:"fileName"`
}

func (pay Payload) Validate() error {
	return validation.ValidateStruct(&pay,
		validation.Field(&pay.FileName, validation.Required),
	)
}
