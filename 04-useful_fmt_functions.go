package main

import "fmt"

func main() {
	helloWord := "Hello"
	worldWord := "world"

	fmt.Println(helloWord, worldWord)

	name := "Anthony"
	age := 22

	// %s for strings and %d for integers
	fmt.Printf("Hi, my name is %s. I'm %d years old.\n", name, age)

	// %v for unknown data types
	fmt.Printf("Hi, my name is %v. I'm %v years old.\n", name, age)

	videoGame := "Elden Ring"

	message := fmt.Sprintf("%s is the GOTY!!!", videoGame)
	fmt.Println(message)

	dInt := 23
	dFloat := 9.8
	dString := "obsidian"
	dBool := true

	fmt.Printf("dInt: %T\n", dInt)
	fmt.Printf("dFloat: %T\n", dFloat)
	fmt.Printf("dString: %T\n", dString)
	fmt.Printf("dBool: %T\n", dBool)
}
