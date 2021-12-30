package controllers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/alienate/password-generator/config"
)

func NewToken() (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(config.Secret)
	if err != nil {
		return t, err
	}
	fmt.Println(t)
	return t, err
}

// func ValidateToken(){

// }
func Verify(bearer string) (string, error) {

	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.Secret, nil
	})
	if err != nil {
		return "nil", err
	}
	fmt.Println(token.Claims.(jwt.Claims))
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return "nil", fmt.Errorf("Validation Error")
	}

	return token.Raw, nil
}
