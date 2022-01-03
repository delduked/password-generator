package controllers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/alienate/password-generator/config"
	"gitlab.com/alienate/password-generator/schema"
)

func NewTokenWithUserName(body *schema.SignUp) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"Username": body.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
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

	// Parse the token to check if it's valid
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

	// Decode the second portion of the JWT token for the username
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "nil", fmt.Errorf("could not parse claims")
	}

	// Return the username from claims
	username, ok := claims["Username"].(string)
	if !ok {
		return "nil", fmt.Errorf("no field Usernamein JWT")
	}

	return username, nil
}

func CheckSecret(body *schema.SignUp) error {
	if body.Secret == "n4th4n43l" {
		return nil
	}
	if body.Secret != "n4th4n43l" {
		return fmt.Errorf("incorrect secret")
	}
	return fmt.Errorf("System error")
}
