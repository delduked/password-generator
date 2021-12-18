package controllers

import (
	"crypto/rand"
	"math/big"

	"gitlab.com/alienate/password-generator/constants"
)

//import "gitlab.com/alienate/password-generator/constants"

func All() *big.Int {
	result, _ := rand.Int(rand.Reader, big.NewInt(67))

	return result
}
func Number() *big.Int {
	result, _ := rand.Int(rand.Reader, big.NewInt(10))

	return result
}
func litleLetter() *big.Int {
	result, _ := rand.Int(rand.Reader, big.NewInt(26))

	return result
}
func upperLetter() *big.Int {
	value, _ := rand.Int(rand.Reader, big.NewInt(26))
	result := constants.Alpha[value]

	return result
}
func Special() *big.Int {
	result, _ := rand.Int(rand.Reader, big.NewInt(31))

	return result
}
