package controllers

import (
	"belajar-gin/database"
	"belajar-gin/models"
	"fmt"
)

func CreateUser(email string) {
	db := database.GetDB()
	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New user data:", User)
}

func GetUserWithProduct() {
	db := database.GetDB()
	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user with product", err.Error())
		return
	}

	fmt.Println("User data with product")
	fmt.Printf("%+v", users)
}
