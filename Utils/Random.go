package Utils

import (
	"math/rand"
	"time"
)

const std_charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const charset = std_charset +
	"~!@#$%^&*(){}[]_+=-/`'\\\"><|"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

func randomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return randomString(length, charset)
}

func StandardRandomString(length int) string {
	return randomString(length, std_charset)
}
