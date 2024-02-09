package controllers

import (
	"belajar-gin/database"
	"belajar-gin/models"
	"fmt"
)

func CreateProduct(userID uint, brand string, name string) {
	db := database.GetDB()
	Product := models.Product{
		UserID: userID,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New product data:", Product)
}
