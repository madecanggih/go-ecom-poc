package config

import (
	"architect/saras-go-poc/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Init(dotenvPath ...string) {

	var conn *gorm.DB
	var err error

	if len(dotenvPath) == 1 {
		err = godotenv.Load(dotenvPath[0])
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Panic("Error loading .env file")
	}

	environment := os.Getenv("ENVIRONMENT")
	fmt.Println("Running " + environment)

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("Connection to: " + dbURI)

	if environment == "PRODUCTION" {
		conn, err = gorm.Open("postgres", dbURI)
	} else {
		conn, err = gorm.Open("sqlite3", ":memory:")
	}

	if err != nil {
		log.Panic(err.Error())
	}

	conn.AutoMigrate(&models.Carts{}, &models.Categories{}, &models.Invoices{}, &models.Products{}, &models.Promos{}, &models.Stores{}, &models.Users{}, &models.Wishilists{})

	DB = conn
}

func Close() {
	DB.Close()
}
