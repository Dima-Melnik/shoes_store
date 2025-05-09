package db

import (
	"log"
	"os"

	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadPassword() string {
	if err := godotenv.Load(); err != nil {
		utils.SendLogrusFatal("env", "loading", err, "Error loading DB password from .env")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		logrus.WithFields(logrus.Fields{
			"component": "env",
			"action":    "getting",
		}).Fatal("Error getting DB password from .env")
	} else {
		log.Println("DB password successfully loaded")
	}

	return dbPassword
}
