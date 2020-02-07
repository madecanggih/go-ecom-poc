package main

import (
	"architect/saras-go-poc/models"
	"architect/saras-go-poc/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	models.InitDB()
	defer models.CloseDB()

	e := routes.InitRoutes()

	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
