package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	ContextKeyRequestID = "request_id"
	ContextKeyUser      = "auth_user"
)

type ValidationItem struct {
	Field   string `json:"field"`
	Rule    string `json:"rule"`
	Message string `json:"message"`
	Value   any    `json:"value,omitempty"`
}

func ParsePositiveInt(raw string, field string) (int, error) {
	value, err := strconv.Atoi(raw)
	if err != nil || value <= 0 {
		return 0, BadRequest("invalid positive integer", err, ValidationItem{
			Field:   field,
			Rule:    "positive_integer",
			Message: "must be a positive integer",
			Value:   raw,
		})
	}
	return value, nil
}

func NormalizeValidationError(err error) any {
	if err == nil {
		return nil
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		return []ValidationItem{{
			Field:   "body",
			Rule:    "json",
			Message: fmt.Sprintf("invalid JSON syntax at offset %d", syntaxErr.Offset),
		}}
	}

	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		return []ValidationItem{{
			Field:   typeErr.Field,
			Rule:    "type",
			Message: fmt.Sprintf("field %s expects %s", typeErr.Field, typeErr.Type.String()),
			Value:   typeErr.Value,
		}}
	}

	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		items := make([]ValidationItem, 0, len(validationErrs))
		for _, fieldErr := range validationErrs {
			items = append(items, ValidationItem{
				Field:   jsonFieldName(fieldErr.Field()),
				Rule:    fieldErr.Tag(),
				Message: validationMessage(fieldErr),
				Value:   fieldErr.Value(),
			})
		}
		return items
	}

	return []ValidationItem{{
		Field:   "request",
		Rule:    "invalid",
		Message: err.Error(),
	}}
}

func validationMessage(fieldErr validator.FieldError) string {
	field := jsonFieldName(fieldErr.Field())
	switch fieldErr.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, fieldErr.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, fieldErr.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, fieldErr.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, fieldErr.Param())
	case "min":
		return fmt.Sprintf("%s minimum length is %s", field, fieldErr.Param())
	case "max":
		return fmt.Sprintf("%s maximum length is %s", field, fieldErr.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", field, fieldErr.Param())
	default:
		return fmt.Sprintf("%s failed validation rule %s", field, fieldErr.Tag())
	}
}

func jsonFieldName(field string) string {
	if field == "" {
		return field
	}

	var b strings.Builder
	for i, r := range field {
		if i > 0 && r >= 'A' && r <= 'Z' {
			b.WriteByte('_')
		}
		b.WriteRune(r)
	}
	return strings.ToLower(b.String())
}
