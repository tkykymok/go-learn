package handlers

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var lock = &sync.Mutex{}

var validateInstance *validator.Validate

func getValidateInstance() *validator.Validate {
	if validateInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		validateInstance = validator.New()
	}
	return validateInstance
}
