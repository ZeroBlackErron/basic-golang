package main

import "fmt"

func main() {
	// (:) of (:=) is used to define a new variable
	// (=) is used to assign a value to an already defined variable
	x := 50
	y := 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	result := x + y
	fmt.Println("Adition:", result)

	result = x - y
	fmt.Println("Subtraction:", result)

	result = x * y
	fmt.Println("Multiplication:", result)

	result = x / y
	fmt.Println("Division:", result)

	result = x % y
	fmt.Println("Module:", result)

	x++
	fmt.Println("Increment:", x)

	y--
	fmt.Println("Decrement:", y)

	// Challenge
	// TODO: Calculate the area of a rectangle, circle and trapezoid

}
