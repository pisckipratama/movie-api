package middleware

import (
	"github.com/asnur/movie-api/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Union type
type Input interface {
	model.Movie | model.User
}

// ErrorResponse is a struct to format error message
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// ValidateStruct is a function to validate struct Data
func ValidateStruct[T Input](data T) []*ErrorResponse {
	var validate = validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

// ValidateField is a middleware to validate field
func ValidateField[T Input]() func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		data := new(T)

		if err := ctx.BodyParser(data); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})

		}

		errors := ValidateStruct(*data)
		if errors != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(errors)

		}

		//Return Next Function
		return ctx.Next()
	}
}
