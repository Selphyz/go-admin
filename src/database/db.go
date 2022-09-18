package database

import (
	"go-admin/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/admin_dashboard"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to DB")
	}
}
func AutoMigrate() {
	err := DB.AutoMigrate(models.User{}, models.Product{}, models.Link{}, models.Order{}, models.OrderItem{})
	if err != nil {
		return
	}
}
