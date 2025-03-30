package Utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mahdic200/weava/Config"
)

type JWTClaims struct {
	Entity_id       int       `json:"entity_id"`
	Table_name      string    `json:"table_name"`
	Expiration_date time.Time `json:"expiration_date"`
	jwt.RegisteredClaims
}

func CreateToken(entity_id int64, table_name string, remember_me bool) (string, time.Time, error) {
	months := 1
	if remember_me {
		months = 6
	}

	expire_time := time.Now().AddDate(0, months, 0)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"entity_id":       entity_id,
		"table_name":      table_name,
		"expiration_date": expire_time,
	})

	tokenString, err := token.SignedString([]byte(Config.JWT_KEY))

	if err != nil {
		return "", time.Time{}, err
	}
	return string(tokenString), expire_time, nil
}

func VerifyToken(tokenString string) (int64, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Config.JWT_KEY), nil
	})

	if err != nil {
		return 0, "", fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token")
	}

	if claims.Expiration_date.Unix() < time.Now().Unix() {
		return 0, "", fmt.Errorf("expired token")
	}

	return int64(claims.Entity_id), claims.Table_name, nil
}
