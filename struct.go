package go_validator

type StructRule struct {
	structPtr interface{}
	validations []Validatable
}

func (i *StructRule) Validate() []error {
	allErrors := make([]error, 0)
	for _, v := range i.validations {
		if errs := v.Validate(); errs != nil {
			allErrors = append(allErrors, errs...)
		}
	}
	if len(allErrors) > 0 {
		return allErrors
	}
	return nil
}


func Struct(structPtr interface{}, validations ...Validatable) *StructRule {
	v := make([]Validatable, len(validations))
	copy(v, validations)
	return &StructRule{structPtr, v}
}
