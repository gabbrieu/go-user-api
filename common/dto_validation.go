package common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Errors []FieldError `json:"errors"`
}

func (e *ValidationError) Error() string {
	b, _ := json.Marshal(e)
	return fmt.Sprintf("validation failed: %s", string(b))
}

var validate = func() *validator.Validate {
	v := validator.New()

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		tag := field.Tag.Get("json")
		if tag == "-" {
			return ""
		}
		name := strings.Split(tag, ",")[0]
		if name == "" {
			return field.Name
		}
		return name
	})

	return v
}()

func humanize(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "field is required"
	case "email":
		return "must be a valid email"
	case "min":
		return fmt.Sprintf("minimum is %s", fe.Param())
	case "max":
		return fmt.Sprintf("maximum is %s", fe.Param())
	case "len":
		return fmt.Sprintf("length must be %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("must be one of: %s", fe.Param())
	case "gte":
		return fmt.Sprintf("must be ≥ %s", fe.Param())
	case "gt":
		return fmt.Sprintf("must be > %s", fe.Param())
	case "lte":
		return fmt.Sprintf("must be ≤ %s", fe.Param())
	case "lt":
		return fmt.Sprintf("must be < %s", fe.Param())
	case "uuid4", "uuid":
		return "must be a valid UUID"
	case "url":
		return "must be a valid URL"
	default:
		return fmt.Sprintf("invalid (%s)", fe.Tag())
	}
}

func Validate(dto any) error {
	if dto == nil {
		return &ValidationError{Errors: []FieldError{{Field: "", Message: "nil payload"}}}
	}

	if err := validate.Struct(dto); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var out []FieldError
		for _, fe := range err.(validator.ValidationErrors) {
			out = append(out, FieldError{
				Field:   fe.Field(),
				Message: humanize(fe),
			})
		}
		return &ValidationError{Errors: out}
	}

	return nil
}
