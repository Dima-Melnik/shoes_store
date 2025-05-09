package config

import (
	"os"

	"github.com/Dima-Melnik/shoes_store/internal/utils"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
		User    string `yaml:"user"`
		DBName  string `yaml:"dbname"`
		SSLMode string `yaml:"sslmode"`
	} `yaml:"database"`
}

var Cfg Config

func LoadConfig(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		utils.SendLogrusFatal("config", "reading", err, "Error reading yaml file")
	}

	if err := yaml.Unmarshal(file, &Cfg); err != nil {
		utils.SendLogrusFatal("config", "unmarshalling", err, "Error unmarshalling yaml file")
	}
}
