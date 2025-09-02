package dto

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

const ALLOWED_USERNAME_SYMBOLS = ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_")

func RegisterValidators(v *validator.Validate) {
	v.RegisterValidation("username", UsernameValidation)
}

func UsernameValidation(f validator.FieldLevel) bool {
	s := f.Field().String()

	for _, char := range []rune(s) {
		if !strings.Contains(ALLOWED_USERNAME_SYMBOLS, string(char)) {
			return false
		}
	}
	return true
}
