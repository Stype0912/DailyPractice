package strong_password_checker_ii

import (
	"strings"
	"unicode"
)

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	var hasLower, hasUpper, hasDigit, hasSpecial bool
	for i, letter := range password {
		if i > 0 && password[i] == password[i-1] {
			return false
		}
		if unicode.IsLower(letter) {
			hasLower = true
		} else if unicode.IsUpper(letter) {
			hasUpper = true
		} else if unicode.IsDigit(letter) {
			hasDigit = true
		} else if strings.ContainsRune("!@#$%^&*()-+", letter) {
			hasSpecial = true
		}
	}
	if hasLower && hasUpper && hasDigit && hasSpecial {
		return true
	} else {
		return false
	}
}
