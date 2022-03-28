package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	upCaseSet      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowCaseSet     = "abcdedfghijklmnopqrstuvwxyz"
	numberSet      = "0123456789"
	specialCharSet = "!@#$%&*"
	allSet         = lowCaseSet + upCaseSet + specialCharSet + numberSet
)

var minUpCase int
var minNumber int
var minSpecial int
var passwordLenght int
var password strings.Builder

func init() {
	flag.IntVar(&minUpCase, "u", 1, "The minimu. uppercase letters")
	flag.IntVar(&minSpecial, "s", 1, "The minimum special characters")
	flag.IntVar(&minNumber, "n", 1, "The minimum numbers")
	flag.IntVar(&passwordLenght, "l", 5, "The lenght of the password")
}

func main() {
	rand.Seed(time.Now().Unix())

	generatePassword(minNumber, minSpecial, minUpCase, passwordLenght)
	Password := shuffelPassword(passwordLenght, minSpecial, minUpCase, minNumber)
	fmt.Println(Password)
}

func generatePassword(passwordLenght, minSpecial, minUpCase, minNumber int) string {

	for i := 0; i < minSpecial; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	for i := 0; i < minNumber; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))

	}

	for i := 0; i < minUpCase; i++ {
		random := rand.Intn(len(upCaseSet))
		password.WriteString(string(upCaseSet[random]))
	}

	return password.String()

}

func shuffelPassword(passwordLenght, minSpecial, minUpCase, minNumber int) string {

	remainingLengt := passwordLenght - minNumber - minSpecial - minUpCase
	for i := 0; i < remainingLengt; i++ {
		random := rand.Intn(len(allSet))
		password.WriteString(string(allSet[random]))
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
