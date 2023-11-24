package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName        string
	AppAddress     string
	AppEnv         string
	JWT_SECRET_KEY string
	ENVIRONMENT    string

	MariaDBConfig MariaDBConfig
	BaseAPI       string
}

var config *Config

func Init() {
	err := godotenv.Load("conf/.env")
	if err != nil {
		log.Printf("[Init] error on loading env from file: %+v", err)
	}

	config = &Config{
		AppName:        os.Getenv("APP_NAME"),
		AppAddress:     os.Getenv("APP_ADDRESS"),
		AppEnv:         os.Getenv("ENVIRONMENT"),
		JWT_SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),

		ENVIRONMENT: os.Getenv("ENVIRONMENT"),

		BaseAPI: os.Getenv("BASE_API"),
		MariaDBConfig: MariaDBConfig{
			DBName:   os.Getenv("DB_NAME"),
			Address:  os.Getenv("DB_ADDRESS"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
}

func Get() *Config {
	return config
}
