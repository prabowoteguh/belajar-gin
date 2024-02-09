package main

import (
	"belajar-gin/database"
	"belajar-gin/routers"

	"fmt"
)

func main() {
	database.StartDB()
	var PORT = ":8000"
	routers.InitApiRoutes().Run(PORT)
	fmt.Println("Application is listening on port", PORT)
}
