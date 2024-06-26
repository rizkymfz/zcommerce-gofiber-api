package database

import (
	"fmt"
	"log"

	"github.com/rizkymfz/zcommerce-gofiber-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	host := config.Config("DB_HOST")
	name := config.Config("DB_NAME")
	user := config.Config("DB_USER")
	pwd := config.Config("DB_PASSWORD")
	port := config.Config("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}

	fmt.Println("Connecting to database...")
}
