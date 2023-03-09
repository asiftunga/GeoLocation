package common

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func HTTPResponser(status_code int, status bool, message string, data ...interface{}) fiber.Map {
	responseData := fiber.Map{
		"error":      status,
		"statuscode": status_code,
		"message":    message,
	}
	if len(data) > 0 {
		responseData["data"] = data[0]
	}
	return responseData
}

func Contains(input_text string, search_text string) bool {
	CheckContains := strings.Contains(input_text, search_text)
	return CheckContains
}
