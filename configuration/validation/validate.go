package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	http_errors "github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/errors"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enTranslator := en.New()
		unt := ut.New(enTranslator, enTranslator)
		transl, _ = unt.GetTranslator("en")
		err := en_translation.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidateUserError(
	validation_err error,
) *http_errors.HttpErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return http_errors.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []http_errors.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := http_errors.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return http_errors.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return http_errors.NewBadRequestError("Error trying to convert fields")
	}
}
