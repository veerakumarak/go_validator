package go_validator

import ("strconv")

type boolRule struct {
	value bool
	kind string
	name string
	rules []*Rule
}

func (i *boolRule) add(validation Validation, message string) *boolRule {
	if i.rules == nil {
		i.rules = make([]*Rule, 0)
	}
	i.rules = append(i.rules, &Rule{validation, message})
	return i
}

func (i *boolRule) True()  *boolRule {
	return i.add(func() bool { return i.value == true}, "is not true")
}

func (i *boolRule) False()  *boolRule {
	return i.add(func() bool { return i.value == false}, "is not false")
}

func (i *boolRule) Validate() []error {
	return Validate(i.rules, i.kind, strconv.FormatBool(i.value))
}

func Bool(value bool) *boolRule {
	return &boolRule{value: value, kind: "bool"}
}
