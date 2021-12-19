package controllers

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

func generate(value int, length int) string {

	var password string
	for i := 0; i < length; i++ {
		// the first, and the second character in the string
		password += iterateOverFunctions(value)
	}

	return password
}
func iterateOverFunctions(value int) string {
	switch value {
	case 0000:
		return ""
	case 0001:
		return characters[rand.Intn(len(constants.Lower))]
	case 0010:
		return characters[rand.Intn(len(constants.Upper))]
	case 0011:
		characters = append(constants.Lower, constants.Upper...)
		characters = append(characters, constants.Lower...)
		return characters[rand.Intn(len(characters))]
	case 0100:
		return characters[rand.Intn(len(constants.Numbers))]
	case 0101:
		characters = append(constants.Numbers, constants.Lower...)
		return characters[rand.Intn(len(characters))]
	case 110:
		characters = append(constants.Numbers, constants.Upper...)
		return characters[rand.Intn(len(characters))]
	case 111:
		characters = append(constants.Numbers, constants.Upper...)
		characters = append(characters, constants.Lower...)
		return characters[rand.Intn(len(characters))]
	case 1000:
		return characters[rand.Intn(len(constants.Special))]
	case 1001:
		characters = append(constants.Special, constants.Lower...)
		return characters[rand.Intn(len(characters))]
	case 1010:
		characters = append(constants.Special, constants.Upper...)
		return characters[rand.Intn(len(characters))]
	case 1011:
		characters = append(constants.Special, constants.Upper...)
		characters = append(characters, constants.Lower...)
		return characters[rand.Intn(len(characters))]
	case 1100:
		characters = append(constants.Special, constants.Numbers...)
		return characters[rand.Intn(len(characters))]
	case 1101:
		characters = append(constants.Special, constants.Numbers...)
		characters = append(characters, constants.Lower...)
		return characters[rand.Intn(len(characters))]
	case 1110:
		characters = append(constants.Special, constants.Numbers...)
		characters = append(characters, constants.Upper...)
		return characters[rand.Intn(len(characters))]
	case 1111:
		characters = append(constants.Special, constants.Numbers...)
		characters = append(characters, constants.Upper...)
		characters = append(characters, constants.Lower...)
		return characters[rand.Intn(len(characters))]
	}

	return "*"
}
func checkLength(length int) int {
	if length > 16 {
		length = 16
	}
	return length
}
