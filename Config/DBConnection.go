package Config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, TIMEZONE)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return fmt.Errorf("config, connect : %s", err.Error())
	}

	DB = db
	return nil
}
