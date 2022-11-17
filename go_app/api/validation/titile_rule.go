package validation

import "github.com/go-playground/validator/v10"

func customValidation(fl validator.FieldLevel) bool {
	ret, _ := fl.Field().Interface().(string)
	return ret == "タイトルテスト"
}
