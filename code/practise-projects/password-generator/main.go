package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

func userInput() (int, int, int, int) {

	fmt.Print("Please enter the number of uppercase letters you want in your password: ")
	var upper int
	var err error

	_, err = fmt.Scanf("%d", &upper)

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return 0, 0, 0, 0
	}

	fmt.Print("Please enter the number of lowercase letters you want in your password: ")
	var lower int
	_, err = fmt.Scanf("%d", &lower)

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return 0, 0, 0, 0
	}

	fmt.Print("Please enter the number of special characters you want in your password: ")
	var special int
	_, err = fmt.Scanf("%d", &special)

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return 0, 0, 0, 0
	}

	fmt.Print("Please enter the number of numbers you want in your password: ")
	var numbers int
	_, err = fmt.Scanf("%d", &numbers)

	fmt.Print("\n")

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return 0, 0, 0, 0
	}

	return upper, lower, special, numbers

}

func shuffleString(s string) string {
	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		jBig, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		j := int(jBig.Int64())
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func cryptoRandSpecial(max int) string {

	chars := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	result := make([]byte, max)
	for i := 0; i < max; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[n.Int64()]
	}
	return string(result)
}

func cryptoRandNumber(max int) (string, error) {

	numbers := max
	s := ""

	for i := 0; i < numbers; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(10)))
		if err != nil {
			return "", err
		}
		s += strconv.Itoa(int(n.Int64()))
	}

	return s, nil
}

func cryptoRandUpper(max int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, max)
	for i := 0; i < max; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[n.Int64()]
	}
	return string(result)
}

func cryptoRandLower(max int) string {
	chars := "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, max)
	for i := 0; i < max; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[n.Int64()]
	}
	return string(result)
}

func randomization(upper int, lower int, special int, numbers int) string {

	Number, err := cryptoRandNumber(numbers)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}

	Special := cryptoRandSpecial(special)
	Lower := cryptoRandLower(lower)
	Upper := cryptoRandUpper(upper)

	/* fmt.Println("The Upper is: ", Upper)
	fmt.Println("The Lower is: ", Lower)
	fmt.Println("The Special is: ", Special)
	fmt.Println("The Number is: ", Number) */

	result := Number + Special + Lower + Upper
	result = shuffleString(result)

	return result
}

func main() {

	fmt.Print("\n")
	fmt.Println("Welcome to the Password Generator!")
	fmt.Println("This application will help you generate a secure password based on your preferences.")
	fmt.Println("Please follow the prompts to customize your password.")
	fmt.Print("\n")

	for {
		upper, lower, special, numbers := userInput()
		result := randomization(upper, lower, special, numbers)
		fmt.Println("Generated password:", result)
		fmt.Printf("Generated password length: %d\n", len(result))
		fmt.Println("Generate another one? (y/n):")
		var choice string
		fmt.Scan("%d", &choice)
		if choice != "y" {
			fmt.Println()
			break
		}
	}

	fmt.Println("Good Bye!")
}
