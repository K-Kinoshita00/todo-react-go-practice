package main

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// トークンを発行するコマンド: go run cmd/issuejwt/main.go
func main() {
	claims := jwt.MapClaims{
		"sub": "dev-user",
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	// トークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// トークンを署名
	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}
	fmt.Println(signed)
}
