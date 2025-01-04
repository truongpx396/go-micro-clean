package config

import (
	"log"
	"project/modules/item/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
