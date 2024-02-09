package database

import (
	"belajar-gin/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "root"
	port     = "5432"
	dbname   = "bootcamp_golang"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}
}

func Migrate() {
	StartDB()
	db.Debug().AutoMigrate(models.User{}, models.Product{}, models.Car{}, models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
