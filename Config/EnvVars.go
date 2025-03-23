package Config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVarsStructure struct {
	APP_BASEURL string
}

func GetEnv() (*EnvVarsStructure, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &EnvVarsStructure{
		APP_BASEURL: os.Getenv("APP_BASEURL"),
	}, nil
}
