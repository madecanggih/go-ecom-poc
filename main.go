package main

import (
	"architect/saras-go-poc/config"
	"architect/saras-go-poc/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	config.Init()
	defer config.Close()

	e := routes.Init()

	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
