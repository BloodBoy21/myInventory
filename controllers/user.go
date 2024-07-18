package controllers

import (
	"github.com/gofiber/fiber/v2"
	"myInventory/helpers"
	"myInventory/models"
	"myInventory/sql"
)

var db = sql.GetClient()

func SingUpUser(c *fiber.Ctx) error {
	userData := new(models.UserIn)
	if err := c.BodyParser(userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid JSON",
		})
	}

	user := models.User{
		Email:    userData.Email,
		Password: userData.Password,
	}

	exists := helpers.ExistsUser(user.Email)
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User already exists",
		})
	}

	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not create user",
		})
	}

	user.Password = hashedPassword

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not create user",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    fiber.Map{"userId": user.ID},
	})
}

func LoginUser(c *fiber.Ctx) error {
	var userIn models.UserIn
	if err := c.BodyParser(&userIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid JSON",
		})
	}
	var user models.User
	db.Where("email = ?", userIn.Email).First(&user)
	err := helpers.ComparePassword(user.Password, userIn.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid credentials",
		})
	}
	token, err := helpers.CreateToken(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not create token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    fiber.Map{"token": token},
	})
}
