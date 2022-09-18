package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/src/database"
	"go-admin/src/models"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order
	database.DB.Preload("OrderItems").Find(&orders)
	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.GetTotal()
	}
	return c.JSON(orders)
}
