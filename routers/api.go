package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func InitApiRoutes() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api")
	{
		orders := apiGroup.Group("/orders")
		{
			orders.GET("/", controllers.OrderIndex)
			orders.GET("/:orderId", controllers.OrderShow)
			orders.POST("/", controllers.OrderStore)
			orders.PUT("/:orderId", controllers.OrderUpdate)
			orders.DELETE("/:orderId", controllers.OrderDestroy)
		}
		cars := apiGroup.Group("/cars")
		{
			cars.POST("/", controllers.CreateCar)
			cars.PUT("/:carID", controllers.UpdateCar)
			cars.GET("/:carID", controllers.GetCar)
			cars.GET("/all", controllers.GetAllCars)
			cars.DELETE("/:carID", controllers.DeleteCar)
		}
	}

	return router
}
