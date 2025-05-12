package jwt

import (
	"log"
	"os"

	"github.com/Dima-Melnik/shoes_store/internal/utils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadPassword() string {
	if err := godotenv.Load(); err != nil {
		utils.SendLogrusFatal("env", "loading", err, "Error loading secret key from .env")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		logrus.WithFields(logrus.Fields{
			"component": "env",
			"action":    "getting",
		}).Fatal("Error getting secret key from .env")
	} else {
		log.Println("Secret key successfully loaded")
	}

	return secretKey
}
