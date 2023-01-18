package main

import "fmt"

func checkDatabaseStatus() {
	// Before the current function ends, this instruction will be executed
	defer fmt.Println("Close DB connection...")

	fmt.Println("Open DB connection...")
	fmt.Println("Running queries...")
}

func main() {
	checkDatabaseStatus()

	// From 0 to 10, just print the odd numbers
	fmt.Println("-- Odd numbers [0, 10]--")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skips to the next iteration
		}
		fmt.Println(i)
	}
	fmt.Println("------------------------")

	// From 0 to 100, stop printing numbers above 20
	fmt.Println("---- Top 20 numbers ----")
	for i := 0; i <= 100; i++ {
		if i > 20 {
			break // The loop ends here
		}
		fmt.Println(i)
	}
	fmt.Println("------------------------")
}
