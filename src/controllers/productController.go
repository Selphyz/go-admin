package controllers

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-admin/src/database"
	"go-admin/src/models"
	"sort"
	"strconv"
	"strings"
	"time"
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
	go database.ClearCache("products_frontend", "products_frontend")
	return c.JSON(product)
}
func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{}
	product.Id = uint(id)
	database.DB.Find(&product)
	return c.JSON(product)
}
func PutProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{}
	product.Id = uint(id)
	database.DB.Model(&product).Updates(&product)
	go database.ClearCache("products_frontend", "products_frontend")
	return c.JSON(product)
}
func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{}
	product.Id = uint(id)
	database.DB.Delete(&product)
	go database.ClearCache("products_frontend", "products_frontend")
	return nil
}
func ProductsFrontend(c *fiber.Ctx) error {
	var products []models.Product
	var ctx = context.Background()
	result, err := database.Cache.Get(ctx, "products_frontend").Result()
	if err != nil {
		database.DB.Find(&products)
		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}
		if errKey := database.Cache.Set(ctx, "products_frontend", bytes, 30*time.Minute).Err(); err != nil {
			panic(errKey)
		}
	} else {
		err := json.Unmarshal([]byte(result), &products)
		if err != nil {
			return err
		}
	}
	return c.JSON(products)
}
func ProductsBackend(c *fiber.Ctx) error {
	var products []models.Product
	var ctx = context.Background()
	result, err := database.Cache.Get(ctx, "products_backend").Result()
	if err != nil {
		database.DB.Find(&products)
		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}
		if errKey := database.Cache.Set(ctx, "products_backend", bytes, 30*time.Minute).Err(); err != nil {
			panic(errKey)
		}
	} else {
		err := json.Unmarshal([]byte(result), &products)
		if err != nil {
			return err
		}
	}
	var searchedProducts []models.Product
	if s := c.Query("s"); s != "" {
		lower := strings.ToLower(s)
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), lower) || strings.Contains(strings.ToLower(product.Description), lower) {
				searchedProducts = append(searchedProducts, product)
			}
		}
	} else {
		searchedProducts = products
	}
	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asd" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price < searchedProducts[j].Price
			})
		} else if sortLower == "desc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price > searchedProducts[j].Price
			})
		}
	}
	total := len(searchedProducts)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 9
	var data []models.Product
	if total <= page*perPage && total >= (page-1)*perPage {
		data = searchedProducts[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = searchedProducts[(page-1)*perPage : page*perPage]
	} else {
		data = []models.Product{}
	}
	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}
