package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hi")
        fmt.Println(rnumberIncrement())
        fmt.Println(multliply())
	fmt.Println(passwordGenerator())
}

func numberIncrement(){
	initialNumber := 5
	incrementedNumber := increment(initialNumber)

	fmt.Printf("The incremented number is: %d\n", incrementedNumber)
} 

func multiply () {
	result := multiply(5, 7)
	return a * b
	fmt.Println("Result:", result)
}


// defined a function GenerateRandomPassword generates a random password of the specified length.
func GenerateRandomPassword(length int) string {

	
	// Defined the data set from the password should generate
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*"

	// Create a password by randomly selecting characters from the charset
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

func passwordGenerator() {
	passwordLength := 16 // Change this to your desired password length

	// Generate a random password
	password := GenerateRandomPassword(passwordLength)

	fmt.Println("Generated Password:", password)
}
