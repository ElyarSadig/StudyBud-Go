package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func ValidateName(name string) error {
	if len(name) < 2 {
		return errors.New("name must be at least 2 characters long")
	}
	if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(name) {
		return errors.New("name can only contain letters and spaces")
	}
	return nil
}

func ValidateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	if !regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_.]*$`).MatchString(username) {
		return errors.New("username can only contain letters, digits, underscores, and dots, and must start with a letter")
	}
	return nil
}

func ValidatePassword(password string) error {
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
		return errors.New("password must be at least 8 characters long")
	}
	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func ValidateEmail(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func FormatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	return fmt.Sprintf("%dh %dm", hours, minutes)
}
