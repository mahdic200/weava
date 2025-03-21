package Config

import (
	"os"

	"github.com/joho/godotenv"
)

const APP_BASEURL="APP_BASEURL"
const DB_PORT="DB_PORT"
const DB_HOST="DB_HOST"
const DB_DBNAME="DB_DBNAME"
const DB_USER="DB_USER"
const DB_PASSWORD="DB_PASSWORD"
const VALIDATION_LANG="VALIDATION_LANG"
const JWT_SECRET="JWT_SECRET"

type EnvVarsInterface struct {
    APP_BASEURL string
    DB_PORT string
    DB_HOST string
    DB_DBNAME string
    DB_USER string
    DB_PASSWORD string
    VALIDATION_LANG string
    JWT_SECRET string
}

func GetEnv() EnvVarsInterface {
    godotenv.Load()
    return EnvVarsInterface{
        APP_BASEURL: os.Getenv(APP_BASEURL),
        DB_PORT: os.Getenv(DB_PORT),
        DB_HOST: os.Getenv(DB_HOST),
        DB_DBNAME: os.Getenv(DB_DBNAME),
        DB_USER: os.Getenv(DB_USER),
        DB_PASSWORD: os.Getenv(DB_PASSWORD),
        VALIDATION_LANG: os.Getenv(VALIDATION_LANG),
        JWT_SECRET: os.Getenv(JWT_SECRET),
    }
}
