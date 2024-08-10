package utility

import (
	"github.com/gofiber/fiber/v2"
)

// ParsePayload parses the request body into a map
func ParsePayload(c *fiber.Ctx) (map[string]interface{}, error) {
	payload := make(map[string]interface{})
	if err := c.BodyParser(&payload); err != nil {
		return nil, err
	}
	return payload, nil
}
