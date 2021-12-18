package controllers

import (
	"crypto/rand"
	"math/big"
	"gitlab.com/alienate/password-generator/inter"
	"gitlab.com/alienate/password-generator/constants"
)

//password length
//include speicals
// include numbers
//inlcude lowerfase
//include uppsercase

func GenerateResponse(length int, specials bool, numbers bool, lowerCases bool, upperCases bool) {

	value := assignment(specials, numbers, lowerCases, upperCases)

}

func assignment(specials bool, numbers bool, lowerCases bool, upperCases bool) int {
	var special int
	var number int
	var lowerCase int
	var upperCase int

	if specials {
		special = 1
	} else {
		special = 0
	}

	if numbers {
		number = 3
	} else {
		number = 0
	}

	if lowerCases {
		lowerCase = 5
	} else {
		lowerCase = 0
	}

	if upperCases {
		upperCase = 7
	} else {
		upperCase = 0
	}

	return special + number + lowerCase + upperCase
}
func checkLength(length int) int {
	if length > 16 {
		length = 16
	}
	return length
}

func generate(value int, length int) {
	switch value {
	case 0:
		// if return value is 16 or 0, client wants all options enabled
		var resp string
		lengthOfPassword := checkLength(length)

		for i := 0; i < 16; i++ {
			// the first, and the second character in the string
		}

	case 1:
		// if value is 1, client wants special
	case 3:
		// if value is 3, client wants numbers
	case 4:
		// if value is 4, client wants specials and numbers
	case 5:
		// if value is 5, client wants lowercase
	case 6:
		// if value is 6, cient wants lowercase and specials
	case 7:
		// if value is 7, client wants uppercase
	case 8:
		// if value is 8, client wants uppercase and specials,
	case 9:
		// if value is 9, client wants lowercase, specials and numbers
	case 11:
		// if value is 11, client wants uppercase specials and numbers
	case 16:
		// if return value is 16 or 0, client wants all options enabled
	}
}
func iterateOverFunctions() string {
	whichFunction, _ := rand.Int(rand.Reader, big.NewInt(4))
	switch whichFunction {
	case 1:
		response := constants.Alpha[]
	case 2:

	case 3:

	case 4:

	}
	return response
}

//import "gitlab.com/alienate/password-generator/constants"

func All() *big.Int {
	result, _ := rand.Int(rand.Reader, big.NewInt(67))

	return result
}
func Number() *big.Int {
	result, _ := rand.Int(rand.Reader, big.NewInt(10))

	return result
}
func Letter() *big.Int {
	result, _ := rand.Int(rand.Reader, big.NewInt(26))

	return result
}
func Special() *big.Int {
	result, _ := rand.Int(rand.Reader, big.NewInt(31))

	return result
}
