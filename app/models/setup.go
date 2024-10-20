package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//driver := os.Getenv("DB_DRIVER")
	//dbUser := os.Getenv("DB_USER")
	//dbPass := os.Getenv("DB_PASS")
	//dbName := os.Getenv("DB_NAME")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")

	//dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	dbURI := "dev:dev@(mariadb)/devDB"
	DB, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	DB.AutoMigrate(&User{})
}
