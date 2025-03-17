package Config

import (
	"fmt"

	"github.com/mahdic200/weava/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
    env := GetEnv()
    connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_DBNAME, env.DB_PORT)
    db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
        SkipDefaultTransaction: true,
    })
    
    if err != nil {
        panic("db connection failed")
    }

    DB = db

    db.AutoMigrate(
        &Models.User {},
    )
}
