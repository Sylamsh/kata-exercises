package password

import (
	"errors"
	"strconv"
)

var Errors = map[string]string{
	"MIN_8_CHAR": "Password must be at least 8 characters",
	"MIN_2_NUM":  "The password must contain at least 2 numbers",
	"MIN_1_CAP":  "Password must contain at least one capital letter",
	"MIN_1_SPEC": "Password must contain at least one special character",
}

func ValidatePassword(password string) (bool, error) {
	errs := ""
	if len(password) < 8 {
		errs += Errors["MIN_8_CHAR"]
	}
	if !hasMin2Numbers(password) {
		errs = addNewlineifNotEmpty(errs)
		errs += Errors["MIN_2_NUM"]
	}
	if !hasAtleastOneCapital(password) {
		errs = addNewlineifNotEmpty(errs)
		errs += Errors["MIN_1_CAP"]
	}
	if !hasAtleastOneSpecial(password) {
		errs = addNewlineifNotEmpty(errs)
		errs += Errors["MIN_1_SPEC"]
	}
	if errs != "" {
		return false, errors.New(errs)
	}
	return true, nil
}

// Checks
func hasMin2Numbers(password string) bool {
	numbers := 0
	for _, char := range password {
		if isNumeric(char) {
			numbers++
		}
	}
	return numbers >= 2
}

func hasAtleastOneCapital(errs string) bool {
	for _, char := range errs {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func hasAtleastOneSpecial(errs string) bool {
	for _, char := range errs {
		if (char < 'A' || char > 'z') && !isNumeric(char) {
			return true
		}
	}
	return false
}

// Helpers
func isNumeric(c rune) bool {
	_, err := strconv.ParseFloat(string(c), 64)
	return err == nil
}

func addNewlineifNotEmpty(errs string) string {
	if errs != "" {
		errs += "\n"
	}
	return errs
}
