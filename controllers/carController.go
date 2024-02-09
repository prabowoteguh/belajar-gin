package controllers

import (
	"fmt"
	"net/http"

	"belajar-gin/database"
	"belajar-gin/models"

	"github.com/gin-gonic/gin"
)

type Car struct {
	ID    uint   `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	db := database.GetDB()
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dataCar := models.Car{
		Brand: newCar.Brand,
		Model: newCar.Model,
		Price: newCar.Price,
	}

	err := db.Create(&dataCar).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New product data:", dataCar)

	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": "Data successfully added!",
		"car":    dataCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	var carParam Car

	db := database.GetDB()
	car := models.Car{}

	if err := ctx.ShouldBindJSON(&carParam); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Model(&car).Where("id = ?", carID).Updates(models.Car{
		Brand: carParam.Brand,
		Model: carParam.Model,
		Price: carParam.Price,
	}).Error

	if err != nil {
		fmt.Println("Error updating car data", err)
		return
	}

	fmt.Println("Car data successfully updated!", carParam)

	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": "Data successfully updated!",
		"car":    carParam,
	})
}

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false

	var carData Car

	for i, car := range CarDatas {
		if carData.ID == car.ID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found!",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	db := database.GetDB()
	car := models.Car{}

	err := db.Where("id = ?", carID).Delete(&car).Error

	if err != nil {
		fmt.Println("Error deleting car data", err)
		return
	}

	fmt.Println("Car data successfully deleted!")

	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": "Data successfully deleted!",
	})
}

func GetAllCars(ctx *gin.Context) {
	db := database.GetDB()
	var cars []Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting cars", err.Error())
		return
	}

	fmt.Println("Cars data")
	fmt.Printf("%+v", cars)

	ctx.JSON(http.StatusOK, gin.H{
		"cars": cars,
	})
}
