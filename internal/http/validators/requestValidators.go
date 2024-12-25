package validators

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/schema"
)

var validate = validator.New()

func ValidateQueryParams(r *http.Request, params interface{}) ([]string, error) {

	if err := r.ParseForm(); err != nil {
		return nil, fmt.Errorf("failed to parse form: %w", err)
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(params, r.Form); err != nil {
		return nil, fmt.Errorf("failed to decode params: %w", err)
	}

	if err := validate.Struct(params); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, fmt.Errorf("invalid validation error: %w", err)
		}

		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' validation failed on tag '%s'", err.Field(), err.Tag()))
		}
		return errorMessages, nil
	}

	return nil, nil
}

func CheckQueryParamsValid(w http.ResponseWriter, r *http.Request, queryParams interface{}) bool {
	if errorMessages, err := ValidateQueryParams(r, &queryParams); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return false
	} else if errorMessages != nil {
		http.Error(w, fmt.Sprintf("Validation failed: %s", errorMessages), http.StatusBadRequest)
		return false
	}
	return true
}
