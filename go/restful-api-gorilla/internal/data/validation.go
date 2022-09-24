package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidationError wraps the validators FieldError so we do not expose this
// to out code
type ValidationError struct {
	validator.FieldError
}

// Error returns validation error as string
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

// Validate the item
// for more detail the returned error can be cast into a
// validator.ValidationErrors collection
//
// if ve, ok := err.(validator.ValidationErrors); ok {
//										fmt.Println(ve.Namespace())
//										fmt.Println(ve.Field())
//										fmt.Println(ve.StructNamespace())
//										fmt.Println(ve.Tag())
//										fmt.Println(ve.ActualTag())
//										fmt.Println(ve.Kind())
//										fmt.Println(ve.Type())
//										fmt.Println(ve.Value())
//										fmt.Println(ve.Param())
//										fmt.Println()
//			}
func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i)

	var returnErrs []ValidationError
	if errs != nil {
		verrs := errs.(validator.ValidationErrors)
		if len(verrs) == 0 {
			return nil
		}

		for _, err := range verrs {
			// cast the FieldError into our ValidationError and append to the slice
			ve := ValidationError{err.(validator.FieldError)}
			returnErrs = append(returnErrs, ve)
		}
	}

	return returnErrs
}

// validateSKU
func validateSKU(fl validator.FieldLevel) bool {
	// SKU must be in the format abc-abc-abc
	skuRegexString := "^[a-z]+-[a-z]+-[a-z]+$"
	skuRegex := regexp.MustCompile(skuRegexString)
	return skuRegex.MatchString(fl.Field().String())
}
