package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

func SuccessResponse(ctx *fiber.Ctx, data interface{}, message string, code ...int) error {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	var statusCode  = fiber.StatusOK
	if len(code) > 0 {
		statusCode = code[0]
	}

	return ctx.Status(statusCode).JSON(response)
}

func FailedResponse(ctx *fiber.Ctx, data interface{}, message string, errors []string) error {
	response := Response{
		Success: false,
		Message: message,
		Data:    data,
		Errors:  errors,
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(response)
}
