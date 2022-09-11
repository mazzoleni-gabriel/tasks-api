package validator

import "github.com/go-playground/validator/v10"

var v = validator.New()

// Validate is a method that provides a structure check.
func Validate(entity interface{}) error {
	return v.Struct(entity)
}
