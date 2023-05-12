package phone_number_validator

import "strings"

type PrefixValidator struct {
}

func NewPrefixValidator() *PrefixValidator {
	return &PrefixValidator{}
}

func (validator PrefixValidator) IsValid(number string) bool {
	validPrefixes := []string{"+", "00"}

	for _, prefix := range validPrefixes {
		if strings.HasPrefix(number, prefix) {
			return true
		}
	}

	return false
}
