package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/src/database"
	"go-admin/src/models"
)

func GetAmbassadors(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}
