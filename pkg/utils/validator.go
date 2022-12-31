package utils

import (
	"github.com/go-playground/validator"
)

// Use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}
