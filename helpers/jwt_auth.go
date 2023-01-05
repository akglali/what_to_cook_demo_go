package helpers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func GenerateJWT(userToken string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(168 * time.Hour).Unix(),
		"authorized": true,
		"user":       userToken,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEYJWT")))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(tokenString)

	return tokenString, nil
}

func VerifyJWT(tokenString string, c *gin.Context) (jwt.Claims, error) {
	secret := []byte(os.Getenv("SECRETKEYJWT"))

	// Parse the token and verify its signature
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check that the signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return secret, nil
	})

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Check that the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claims", claims["exp"])
		return claims, nil
	} else {
		fmt.Println("invalid token")
		return nil, errors.New("Invalid Token")
	}
}
