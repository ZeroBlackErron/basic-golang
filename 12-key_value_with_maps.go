package main

import "fmt"

func main() {
	// Map declaration with string keys and integer values
	ages := make(map[string]int)

	ages["Joseph"] = 19
	ages["Camile"] = 18
	ages["Dave"] = 34
	ages["Luna"] = 25

	fmt.Println(ages)

	// Iterating the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Getting a specific value
	josephAge := ages["Joseph"]
	fmt.Println("Joseph age:", josephAge)

	// Since we don't have a 'Carlos' key, it is assigned a zero-value
	carlosAge := ages["Carlos"]
	fmt.Println("Carlos age:", carlosAge)

	// Determining if a key exists in a map
	johnAge, exists := ages["John"]
	fmt.Println("John age:", johnAge, exists)

	camileAge, exists := ages["Camile"]
	fmt.Println("Camile age:", camileAge, exists)
}
