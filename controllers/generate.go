package controllers

import (
	"crypto/rand"
	"math/big"

	"gitlab.com/alienate/password-generator/constants"
	"gitlab.com/alienate/password-generator/schema"
)

func GenerateResponse(req *schema.NewPasswordRequest) string {

	value := assignment(req.Special, req.Number, req.Lower, req.Upper)
	validatedLength := checkLength(req.Length)

	generatedPassword := combinations(value, validatedLength)

	return generatedPassword
}

func assignment(specials bool, numbers bool, lowerCases bool, upperCases bool) int {
	var special int
	var number int
	var lowerCase int
	var upperCase int

	if lowerCases {
		lowerCase = 1
	} else {
		lowerCase = 0
	}

	if upperCases {
		upperCase = 10
	} else {
		upperCase = 0
	}

	if numbers {
		number = 100
	} else {
		number = 0
	}

	if specials {
		special = 1000
	} else {
		special = 0
	}

	return special + number + lowerCase + upperCase
}

func combinations(value int, length int) string {
	var characters string

	switch value {
	case 0:
		return ""
	case 1:
		characters = constants.Lower
		return generateRandomString(length, characters)
	case 10:
		characters = constants.Upper
		return generateRandomString(length, characters)
	case 11:
		characters = constants.Upper + constants.Lower
		return generateRandomString(length, characters)
	case 100:
		characters = constants.Numbers
		return generateRandomString(length, characters)
	case 101:
		characters = constants.Numbers + constants.Lower
		return generateRandomString(length, characters)
	case 110:
		characters = constants.Numbers + constants.Upper
		return generateRandomString(length, characters)
	case 111:
		characters = constants.Numbers + constants.Upper + constants.Lower
		return generateRandomString(length, characters)
	case 1000:
		characters = constants.Special
		return generateRandomString(length, characters)
	case 1001:
		characters = constants.Special + constants.Lower
		return generateRandomString(length, characters)
	case 1010:
		characters = constants.Special + constants.Upper
		return generateRandomString(length, characters)
	case 1011:
		characters = constants.Special + constants.Upper + constants.Lower
		return generateRandomString(length, characters)
	case 1100:
		characters = constants.Special + constants.Numbers
		return generateRandomString(length, characters)
	case 1101:
		characters = constants.Special + constants.Numbers + constants.Lower
		return generateRandomString(length, characters)
	case 1110:
		characters = constants.Special + constants.Numbers + constants.Upper
		return generateRandomString(length, characters)
	case 1111:
		characters = constants.Special + constants.Numbers + constants.Upper + constants.Lower
		return generateRandomString(length, characters)
	}

	return "*"
}

// https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func generateRandomString(passwordLength int, letters string) string {
	ret := make([]byte, passwordLength)
	for i := 0; i < passwordLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "*"
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret)
}

func checkLength(length int) int {
	if length > 32 {
		length = 32
	} else if length < 4 {
		length = 4
	}
	return length
}
