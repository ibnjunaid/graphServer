package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//SecretKey is secret
var (
	SecretKey = []byte("Secret")
)

//GenerateToken generates JWT
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	//Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in generating key")
		return "", nil
	}
	return tokenString, nil
}

//ParseToken is ..
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	}
	return "", err
}
