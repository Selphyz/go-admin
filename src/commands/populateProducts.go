package main

import (
	"github.com/bxcodec/faker/v3"
	"go-admin/src/database"
	"go-admin/src/models"
	"math/rand"
)

func main() {
	database.Connect()
	for i := 0; i < 10; i++ {
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 5),
		}
		database.DB.Create(&product)
	}
}
