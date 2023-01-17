package main

import "fmt"

func main() {
	// 'Go' has only one loop structure

	// Classic 'for' - with the initialization, conditional stop and increment
	fmt.Println("--------- For ---------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("-----------------------")

	// For while - with the conditional stop and the increment
	fmt.Println("------ For while ------")
	firstCounter := 0

	for firstCounter < 10 {
		fmt.Println(firstCounter)
		firstCounter++
	}
	fmt.Println("-----------------------")

	// For forever - it explains by itself
	fmt.Println("----- For forever -----")
	secondCounter := 0

	for {
		fmt.Println(secondCounter)
		secondCounter++
	}
	fmt.Println("-----------------------") // Obviously this line is never reached
}
