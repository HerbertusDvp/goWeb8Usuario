package utils

import (
	"regexp"
	"unicode"
)

var RegexCorreo = regexp.MustCompile("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")

func ValidaPassword(s string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)

	if len(s) >= 6 && len(s) <= 10 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}
