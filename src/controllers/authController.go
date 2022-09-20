package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/src/database"
	"go-admin/src/middleware"
	"go-admin/src/models"
	"strings"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "passwords do not match"})
	}
	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		IsAmbassador: strings.Contains(c.Path(), "/api/ambassador"),
	}
	user.SetPassword(data["password"])
	database.DB.Create(&user)
	return c.JSON(user)
}
func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Invalid credentials"})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		return c.JSON(fiber.Map{"message": "Invalid credentials"})
	}
	isAmbassador := strings.Contains(c.Path(), "/api/ambassador")
	var scope string
	if isAmbassador {
		scope = "ambassador"
	} else {
		scope = "admin"
	}
	if !isAmbassador && user.IsAmbassador {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	token, err := middleware.GenerateJWT(user.Id, scope)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Invalid credentials"})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
func User(c *fiber.Ctx) error {
	id, _ := middleware.GetUserID(c)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	id, _ := middleware.GetUserID(c)
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	user.Id = id
	database.DB.Model(&user).Updates(&user)
	return c.JSON(user)
}
func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "passwords do not match"})
	}
	id, _ := middleware.GetUserID(c)
	user := models.User{}
	user.Id = id
	user.SetPassword(data["password"])
	database.DB.Model(&user).Updates(&user)
	return c.JSON(user)
}
