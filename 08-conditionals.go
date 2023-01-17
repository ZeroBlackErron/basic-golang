package main

import (
	"fmt"
	"log"
	"strconv"
)

func isEven(number int) bool {
	module := number % 2

	return module == 0
}

func isOdd(number int) bool {
	return !isEven(number)
}

func validateUserCredentials(username, password string) bool {
	return username == "admin" && password == "123456"
}

func main() {
	x := 199
	y := 377

	if x == 199 {
		fmt.Println("X es 199")
	} else {
		fmt.Println("X no es 199")
	}

	if x != 911 && y != 733 {
		fmt.Println("All conditions are true")
	}

	if x == 911 || y != 733 {
		fmt.Println("At least one condition is true")
	}

	// Plus: Convert string to number
	number, err := strconv.Atoi("666")

	if err != nil {
		log.Fatal(err) // Print in console
	}

	fmt.Println("Number:", number)

	// Challenge
	// Determine if a number is odd or even and check username and password

	// Odd - Even
	a := 45
	fmt.Printf("%d is even: %t\n", a, isEven(a))
	fmt.Printf("%d is odd: %t\n", a, isOdd(a))

	b := 82
	fmt.Printf("%d is even: %t\n", b, isEven(b))
	fmt.Printf("%d is odd: %t\n", b, isOdd(b))

	// Check username and password
	u := "hacker-man"
	p := "098765"
	fmt.Printf("Access for %s: %t\n", u, validateUserCredentials(u, p))

	u = "admin"
	p = "123456"
	fmt.Printf("Access for %s: %t\n", u, validateUserCredentials(u, p))
}
