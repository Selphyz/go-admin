package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/src/database"
	"go-admin/src/models"
	"strconv"
)

func Link(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var links []models.Link
	database.DB.Where("user_id = ?", id).Find(&links)
	for i, link := range links {
		var orders []models.Order
		database.DB.Where("code = ? and complete = true", link.Code).Find(&orders)
		links[i].Orders = orders
	}
	return c.JSON(links)
}
