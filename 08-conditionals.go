package main

import (
	"fmt"
	"log"
	"strconv"
)

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
	// TODO: Determine if a number is odd or even and check username and password
}
