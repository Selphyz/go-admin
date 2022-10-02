package controllers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"go-admin/src/database"
	"go-admin/src/middleware"
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

type CreateLinkRequest struct {
	Products []int
}

func CreateLink(c *fiber.Ctx) error {
	var request CreateLinkRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	id, _ := middleware.GetUserID(c)
	link := models.Link{
		UserId: id,
		Code:   faker.Username(),
	}
	for _, productId := range request.Products {
		product := models.Product{}
		product.Id = uint(productId)
		link.Products = append(link.Products, product)
	}
	database.DB.Create(&link)
	return c.JSON(link)
}
