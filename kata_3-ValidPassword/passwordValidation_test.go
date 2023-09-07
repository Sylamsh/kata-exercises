package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword_8CharacterLongFalse(t *testing.T) {
	isValid, err := ValidatePassword("123@5A")
	assert.Equal(t, isValid, false)
	assert.Equal(t, err.Error(), "Password must be at least 8 characters")
}

func TestValidatePassword_2NumbersFalse(t *testing.T) {
	isValid, err := ValidatePassword("Abcde@fgh")
	assert.Equal(t, isValid, false)
	if err != nil {
		assert.Equal(t, err.Error(), "The password must contain at least 2 numbers")
	}
}

func TestValidatePassword_1CapitalFalse(t *testing.T) {
	isValid, err := ValidatePassword("abcde$e21")
	assert.Equal(t, isValid, false)
	if err != nil {
		assert.Equal(t, err.Error(), "Password must contain at least one capital letter")
	}
}

func TestValidatePassword_1SpecialCharFalse(t *testing.T) {
	isValid, err := ValidatePassword("Abcdese21")
	assert.Equal(t, isValid, false)
	if err != nil {
		assert.Equal(t, err.Error(), "Password must contain at least one special character")
	}
}

func TestValidatePassword_MulitpleErrors(t *testing.T) {
	isValid, err := ValidatePassword("abc")
	assert.Equal(t, isValid, false)
	if err != nil {
		assert.Contains(t, err.Error(), "The password must contain at least 2 numbers")
		assert.Contains(t, err.Error(), "Password must be at least 8 characters")
		assert.Contains(t, err.Error(), "Password must contain at least one capital letter")
		assert.Contains(t, err.Error(), "Password must contain at least one special character")
	}
}

func TestValidatePassword_Pass(t *testing.T) {
	isValid, err := ValidatePassword("Hello@123")
	assert.Equal(t, isValid, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
