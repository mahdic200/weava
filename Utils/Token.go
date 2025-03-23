package Utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mahdic200/weava/Config"
)

type JWTClaims struct {
	jwt.Claims
	entity_id       int
	table_name      string
	expiration_date time.Time
}

func CreateToken(entity_id int64, table_name string, remember_me bool) (string, error) {
	months := 1
	if remember_me {
		months = 6
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"entity_id":       entity_id,
		"table_name":      table_name,
		"expiration_date": time.Now().AddDate(0, months, 0).Unix(),
	})

	tokenString, err := token.SignedString([]byte(Config.JWT_KEY))

	if err != nil {
		return "", err
	}
	return string(tokenString), nil
}

func VerifyToken(tokenString string) (int64, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Config.JWT_KEY), nil
	})

	if err != nil {
		return 0, "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(JWTClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token")
	}

	if claims.expiration_date.Unix() < time.Now().Unix() {
		return 0, "", fmt.Errorf("expired token")
	}

	return int64(claims.entity_id), claims.table_name, nil
}
