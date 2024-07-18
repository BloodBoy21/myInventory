package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"myInventory/helpers"
)

func IsAuthenticated(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	auth = auth[len("Bearer "):]
	data, err := helpers.VerifyToken(auth)
	if err != nil {
		log.Printf("Error: %s", err)
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	user, err := helpers.GetUserById(int32(data["userId"].(float64)))
	if err != nil {
		log.Printf("Error: %s", err)
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	c.Locals("user", user)
	return c.Next()
}
