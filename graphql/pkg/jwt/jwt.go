package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var SecretKey = []byte("secret")

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
