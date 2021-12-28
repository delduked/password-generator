package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secret = []byte("unicorns")

func CheckForJwt() (string, error) {
	//authToken := c.Cookies("authToken")
	// if authToken == "" {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2023, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "nil", err
	}
	// setCookie := &fiber.Cookie{
	// 	Name:  "authToken",
	// 	Value: tokenString,
	// }
	// c.Cookie(setCookie)

	return tokenString, err
	//c.Next()
	// }

	// _, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf(("Invalid Signing Method"))
	// 	}
	// 	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
	// 		return nil, fmt.Errorf(("Expired token"))
	// 	}

	// 	return secret, nil
	// })
	// if err != nil {
	// 	return "invalid", nil
	// }
	// return "valid", nil

}
