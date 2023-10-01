package database

import (
	"fmt"
	"log"

	"github.com/beruang43221/assignment2-015/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "12345"
	dbPort   = "5432"
	dbname   = "pemesanan"
	db * gorm.DB

	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",host,user,password,dbname,dbPort)

	db,err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connecting Database: ",err)
	}

	fmt.Println("Connecting Database")

	db.Debug().AutoMigrate(models.Order{},models.Item{})
}

func GetDB() *gorm.DB {
	return db
}