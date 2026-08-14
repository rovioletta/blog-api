package validator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func DecodeAndValidateRequest[T any](r *http.Request, v *validator.Validate) (body *T, err error) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		return nil, fmt.Errorf("failed to parse body: %w", err)
	}

	if err := v.Struct(body); err != nil {
		return nil, fmt.Errorf("failed to validate body: %w", err)
	}

	return body, nil
}