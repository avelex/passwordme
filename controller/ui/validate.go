package ui

import (
	"github.com/go-playground/validator"
)

func isValidUrl(str string) bool {
	return validator.New().Var(str, "hostname") == nil
}
