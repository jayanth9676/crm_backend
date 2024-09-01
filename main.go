package main

import (
	"your_project/routes"
	"your_project/config"
)

func main() {
	config.ConnectDB()
	router := routes.SetupRouter()
	router.Run(":8080")
}