package Config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVarsStructure struct {
	APP_BASEURL string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	TIMEZONE    string
}

func GetEnv() (*EnvVarsStructure, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &EnvVarsStructure{
		APP_BASEURL: os.Getenv("APP_BASEURL"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		TIMEZONE: os.Getenv("TIMEZONE"),
	}, nil
}
