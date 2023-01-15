package main

import "fmt"

func main() {
	// Constants declaration
	const pi float64 = 3.1416 // Explicit
	const pi2 = 3.1417        // Implicit

	fmt.Println("------- Constants declaration -------")
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	fmt.Println("-------------------------------------")

	// Variables declaration
	base := 12
	var height int = 14
	var area int

	fmt.Println("------- Variables declaration -------")
	fmt.Println(base, height, area)
	fmt.Println("-------------------------------------")

	// Zero values (not null)
	// golang assign default values for uninitialized variables
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("------------ Zero values ------------")
	fmt.Println(a, b, c, d)
	fmt.Println("-------------------------------------")
}
