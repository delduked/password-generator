package controllers

//password length
//include speicals
// include numbers
//inlcude lowerfase
//include uppsercase
import (
	"math/rand"

	"gitlab.com/alienate/password-generator/constants"
	"gitlab.com/alienate/password-generator/inter"
)

var characters = []string{}

func GenerateResponse(req *inter.Request) string {

	value := assignment(req.Special, req.Number, req.Lower, req.Upper)
	validatedLength := checkLength(req.Length)

	generatedPasswordForClient := generate(value, validatedLength)

	return generatedPasswordForClient
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

func generate(value int, length int) string {

	var password string
	for i := 0; i < length; i++ {
		// the first, and the second character in the string
		password += iterateOverFunctions(value)
	}

	return password
}
func iterateOverFunctions(value int) string {
	// if return value is 16 or 0, client wants all options enabled
	if value == 0 || value == 16 {
		characters = append(constants.Lower, constants.Upper...)
		characters = append(characters, constants.Numbers...)
		characters = append(characters, constants.Special...)

		//loopOverTheAmountOfCharacters := rand.Intn(characters.Length)
		return characters[rand.Intn(len(characters)-1)]
	}

	// if value is 1, client wants special
	if value == 1 {
		characters = constants.Special
		return characters[rand.Intn(len(characters)-1)]
	}

	// if value is 3, client wants numbers
	if value == 3 {
		characters = constants.Numbers
		return characters[rand.Intn(len(characters)-1)]
	}

	// if value is 4, client wants specials and numbers
	if value == 4 {
		characters = append(constants.Special, constants.Numbers...)
		return characters[rand.Intn(len(characters)-1)]
	}

	// if value is 5, client wants lowercase
	if value == 5 {
		characters = constants.Lower
		return characters[rand.Intn(len(characters)-1)]
	}

	// if value is 6, cient wants lowercase and specials
	if value == 6 {
		characters = constants.Lower
		characters = append(characters, constants.Special...)
		return characters[rand.Intn(len(characters)-1)]
	}

	// if value is 7, client wants uppercase
	if value == 7 {
		characters = constants.Upper
		return characters[rand.Intn(len(characters)-1)]
	}

	// if value is 8, client wants uppercase and specials
	if value == 8 {
		characters = constants.Upper
		characters = append(characters, constants.Special...)
		return characters[rand.Intn(len(characters)-1)]
	}
	//if value is 9, client wants lowercase, specials and numbers
	if value == 9 {
		characters = constants.Lower
		characters = append(characters, constants.Special...)
		characters = append(characters, constants.Numbers...)
		return characters[rand.Intn(len(characters)-1)]
	}

	// if value is 11, client wants uppercase specials and numbers
	if value == 9 {
		characters = constants.Upper
		characters = append(characters, constants.Special...)
		characters = append(characters, constants.Numbers...)
		return characters[rand.Intn(len(characters)-1)]
	}
	// if return value is 16 or 0, client wants all options enabled
	return "asdf"
}
func checkLength(length int) int {
	if length > 16 {
		length = 16
	}
	return length
}
