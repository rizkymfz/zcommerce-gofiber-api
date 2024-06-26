package migrations

import (
	"fmt"
	"log"

	"github.com/rizkymfz/zcommerce-gofiber-api/database"
	"github.com/rizkymfz/zcommerce-gofiber-api/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}
