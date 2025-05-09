package db

import (
	"fmt"
	"log"

	"github.com/Dima-Melnik/shoes_store/config"
	"github.com/Dima-Melnik/shoes_store/internal/models"
	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error

	dbPassword := LoadPassword()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.User,
		dbPassword,
		config.Cfg.Database.DBName,
		config.Cfg.Database.SSLMode)

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		utils.SendLogrusFatal("db", "opening", err, "Error opening DB")
	}

	err = DB.AutoMigrate(&models.Attribute{}, &models.Category{}, &models.Product{}, &models.ProductAttributes{}, &models.User{})
	if err != nil {
		utils.SendLogrusFatal("db", "migration", err, "Error migration models to DB")
	}
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error init sqlDB for closing DB")
	}

	if err := sqlDB.Close(); err != nil {
		utils.SendLogrusFatal("db", "closing", err, "Error closing DB")
	}
}
