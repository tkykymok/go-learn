package validation

import (
	"github.com/go-playground/validator/v10"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

var validateInstance *validator.Validate

func GetValidateInstance() *validator.Validate {
	if validateInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		validateInstance = validator.New()

		err := validateInstance.RegisterValidation("title-custom", customValidation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return validateInstance
}
