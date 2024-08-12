package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func ValidateName(name string) (bool, string) {
	if len(name) < 2 {
		return false, "name must be at least 2 characters long"
	}
	if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(name) {
		return false, "name can only contain letters and spaces"
	}
	return true, ""
}

func ValidateUsername(username string) (bool, string) {
	if len(username) < 3 {
		return false, "username must be at least 3 characters long"
	}
	if !regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_.]*$`).MatchString(username) {
		return false, "username can only contain letters, digits, underscores, and dots, and must start with a letter"
	}
	return true, ""
}

func ValidatePassword(password string) (bool, string) {
	var (
		hasMinLen    = false
		hasUpper     = false
		hasLower     = false
		hasDigit     = false
		hasSpecial   = false
		specialChars = "!@#$%^&*()-_=+[]{}|;:'\",.<>?/~`"
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case strings.ContainsRune(specialChars, char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return false, "password must be at least 8 characters long"
	}
	if !hasUpper {
		return false, "password must contain at least one uppercase letter"
	}
	if !hasLower {
		return false, "password must contain at least one lowercase letter"
	}
	if !hasDigit {
		return false, "password must contain at least one digit"
	}
	if !hasSpecial {
		return false, "password must contain at least one special character"
	}

	return true, ""
}

func ValidateEmail(email string) (bool, string) {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return false, "invalid email format"
	}
	return true, ""
}