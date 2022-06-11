package validators

import (
	"github.com/EdricT07/workhours/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(workhours models.User) []*models.ErrorResponse {
	var errors []*models.ErrorResponse
	err := validate.Struct(workhours)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)

		}
	}

	return errors
}
