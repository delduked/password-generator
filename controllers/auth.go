package controllers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/alienate/password-generator/config"
	"gitlab.com/alienate/password-generator/schema"
)

func GenerateNewToken() (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(config.Secret)
	if err != nil {
		return t, err
	}
	return t, err
}

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

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return "nil", fmt.Errorf("Validation Error")
	}

	return token.Raw, nil
}

func CheckCredentials(body *schema.UserAccount) error {

	if body.Username == "nate" && body.Password == "n4t4hn43l" {
		return nil
	}
	if body.Username != "nate" || body.Password != "n4t4hn43l" {
		return fmt.Errorf("incorrect username or password")
	}
	return fmt.Errorf("System error")
}
