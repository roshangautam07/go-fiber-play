package validation

import (
	"api/go/customValidation"
	"api/go/helper/utility"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ValidateUser(c *fiber.Ctx) error {
	payload, _ := utility.ParsePayload(c)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Invalid request payload",
	// 	})
	// }
	errors := customValidation.ValidatePayload(payload)

	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": errors})
	} else {
		fmt.Println("Validation passed!")
	}
	return c.Next()
}
