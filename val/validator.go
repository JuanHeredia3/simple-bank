package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must be between %d and %d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(username string) error {
	if err := ValidateString(username, 3, 20); err != nil {
		return err
	}
	if !isValidUsername(username) {
		return fmt.Errorf("username can only contain alphanumeric characters and underscores")
	}
	return nil
}

func ValidateFullName(fullName string) error {
	if err := ValidateString(fullName, 2, 100); err != nil {
		return err
	}
	if !isValidFullName(fullName) {
		return fmt.Errorf("full name can only contain letters and spaces")
	}
	return nil
}

func ValidatePassword(password string) error {
	return ValidateString(password, 6, 50)
}

func ValidateEmail(email string) error {
	if err := ValidateString(email, 3, 100); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email format: %w", err)
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
