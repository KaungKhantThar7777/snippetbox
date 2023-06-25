package validator

import (
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}

}

func (v *Validator) CheckField(invalid bool, key, message string) {
	if invalid {
		v.AddFieldError(key, message)
	}
}

func (v *Validator) IsEmpty(val string) bool {
	return strings.TrimSpace(val) == ""
}

func (v *Validator) MaxChars(val string, n int) bool {
	return utf8.RuneCountInString(val) > n
}

func PermittedInt(val int, permittedValues ...int) bool {
	for i := range permittedValues {
		if val == permittedValues[i] {
			return true
		}
	}

	return false
}
