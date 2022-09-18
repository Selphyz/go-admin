package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/src/database"
	"go-admin/src/models"
	"strconv"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Find(&products)
	return c.JSON(products)
}
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	database.DB.Create(&product)
	return c.JSON(product)
}
func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{
		Id: uint(id),
	}
	database.DB.Find(&product)
	return c.JSON(product)
}
func PutProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{
		Id: uint(id),
	}
	database.DB.Model(&product).Updates(&product)
	return c.JSON(product)
}
func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{
		Id: uint(id),
	}
	database.DB.Delete(&product)
	return nil
}
