package presenter

import (
	"app/configs"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
	Message     string
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func ValidationErrorResponse(err error) *fiber.Map {
	var errors []*ValidationError
	fieldNamesMap := configs.GetFieldNames()
	for _, err := range err.(validator.ValidationErrors) {
		var temp ValidationError
		temp.FailedField = err.StructField()
		temp.Tag = err.Tag()
		temp.Value = err.Param()
		temp.Message = fieldNamesMap[strings.ToLower(err.StructField())] + "は必須です。" // TODO メッセージの設定を共通化
		errors = append(errors, &temp)
	}

	return &fiber.Map{
		"errors": errors,
	}
}
