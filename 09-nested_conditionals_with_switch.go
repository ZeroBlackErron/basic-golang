package main

import "fmt"

func main() {
	// Firs way
	module := 42 % 2

	switch module {
	case 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}

	// Second way
	switch progress := 100; progress {
	case 100:
		fmt.Println("Complete!")
	default:
		fmt.Println("Not Complete!")
	}

	// Without expression
	testScore := 20

	switch {
	case testScore > 90:
		fmt.Println("Rank A")
	case testScore <= 90 && testScore > 70:
		fmt.Println("Rank B")
	case testScore <= 70 && testScore > 50:
		fmt.Println("Rank C")
	case testScore <= 50 && testScore > 30:
		fmt.Println("Rank D")
	case testScore <= 30 && testScore >= 10:
		fmt.Println("Rank E")
	default:
		fmt.Println("Unranked")
	}
}
