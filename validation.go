package go_validator

import (
	"errors"
	// "reflect"
)

type Validation func() bool

type Validatable interface {
	Validate() []error
	//GetValidations() []Validation
}

type Rule struct{
	validation Validation
	message string
}

// func Validate(validations []Validation) error {
// 	for _, v := range validations {
// 		if !v() {
// 			return errors.New("failed")
// 		}
// 	}
// 	return nil
// }

func Validate(rules []*Rule, kind string, value string) []error {
	errs := make([]error, 0)
	for _, v := range rules {
		if !v.validation() {
			errs = append(errs, errors.New("Error: Type-" + kind + ", Value-" + value + ", " + v.message))
		}
	}
	if len(errs) > 0 {
		return errs;
	}
	return nil
}
