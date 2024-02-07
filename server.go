package main

import (
	"belajar-gin/routers"
	"fmt"
)

func main() {
	var PORT = ":8000"
	routers.StartServer().Run(PORT)

	fmt.Println("Application is listening on port", PORT)
}
