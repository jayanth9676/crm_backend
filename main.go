package main

import (
	"routes"
	"config"
)

func main() {
	config.ConnectDB()
	router := routes.SetupRouter()
	router.Run(":8080")
}