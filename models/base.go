package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //need to be here
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.DB

func init() {

	err := godotenv.Load()
	if err != nil {
		print(err)
	}

	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbUser, dbName, dbPassword)
	conn, err := gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{})
}

// GetDB get the db
func GetDB() *gorm.DB {
	return db
}
