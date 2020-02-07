package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

type (
	DBInterface interface {
		SelectAll() []Users
		SelectById(id int) Users
	}

	DBImplementation struct {
		db *gorm.DB
	}
)

var DB *gorm.DB

func InitDB(dotenvPath ...string) {
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

	conn.AutoMigrate(&Carts{}, &Categories{}, &Invoices{}, &Products{}, &Promos{}, &Stores{}, &Users{}, &Wishilists{})

	DB = conn
}

func CloseDB() {
	DB.Close()
}

func NewDB(db *gorm.DB) *DBImplementation {
	return &DBImplementation{db}
}

func (d *DBImplementation) SelectById(id int) Users {
	users := Users{}
	d.db.Find(&users, id)

	return users
}

func (d *DBImplementation) SelectAll() []Users {
	users := []Users{}
	d.db.Find(&users)

	return users
}
