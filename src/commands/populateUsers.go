package commands

import (
	"github.com/bxcodec/faker/v3"
	"go-admin/src/database"
	"go-admin/src/models"
)

func main() {
	database.Connect()
	for i := 0; i < 10; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}
		ambassador.SetPassword("1234")
		database.DB.Create(&ambassador)
	}
}
