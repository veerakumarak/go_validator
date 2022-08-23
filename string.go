package go_validator

import ("strconv")

type stringRule struct {
	value string
	kind string
	name string
	rules []*Rule
}

func (i *stringRule) add(validation Validation, message string) *stringRule {
	if i.rules == nil {
		i.rules = make([]*Rule, 0)
	}
	i.rules = append(i.rules, &Rule{validation, message})
	return i
}


func (i *stringRule) Empty()  *stringRule {
	// "String val is not empty"
	return i.add(func() bool { return i.value == ""}, "is not empty")
}

func (i *stringRule) NotEmpty()  *stringRule {
	return i.add(func() bool { return i.value != ""}, "is empty")
}

func (i *stringRule) Min(length int)  *stringRule {
	return i.add(func() bool { return len(i.value) >= length}, "has length less than " + strconv.Itoa(length))
}

func (i *stringRule) Max(length int)  *stringRule {
	return i.add(func() bool { return len(i.value) <= length}, "has length more than " + strconv.Itoa(length))
}

func (i *stringRule) Validate() []error {
	return Validate(i.rules, i.kind, i.value)
}

func String(value string) *stringRule {
	return &stringRule{value: value, kind: "string"}
}
