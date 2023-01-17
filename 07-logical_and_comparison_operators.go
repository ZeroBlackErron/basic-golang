package main

import "fmt"

func main() {
	// Comparison operators
	x := 45
	y := 29

	fmt.Printf("%d is equal to %d: ", x, y)
	fmt.Println(x == y)

	fmt.Printf("%d is different from %d: ", x, y)
	fmt.Println(x != y)

	fmt.Printf("%d is less than %d: ", x, y)
	fmt.Println(x < y)

	fmt.Printf("%d is greater to %d: ", x, y)
	fmt.Println(x > y)

	fmt.Printf("%d is less than or equal to %d: ", x, y)
	fmt.Println(x <= y)

	fmt.Printf("%d is greater than or equal to %d: ", x, y)
	fmt.Println(x >= y)

	// Logical operators
	a := 12
	b := 34
	c := 56

	fmt.Printf("%d is different from %d and less than %d: ", b, c, a)
	fmt.Println(b != c && c < a)

	fmt.Printf("%d is different from %d or less than %d: ", b, c, a)
	fmt.Println(b != c || c < a)
}
