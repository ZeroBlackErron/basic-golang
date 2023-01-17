package main

import "fmt"

// Function declaration
func printHelloWorld() {
	fmt.Println("Hello world")
}

// Function declaration with param
func printMessage(message string) {
	fmt.Println(message)
}

// Function declaration with params of the same data type
func printParams(a, b int, c string) {
	fmt.Println(a, b, c)
}

// Function declaration with a single return value
func multiply(x int, y int) int { // return data type must be specified
	return x * y
}

// Function declaration with multiple return values
func getNameAndAge() (name string, age int) {
	return "Anthony", 22
}

func calculateRectangleArea(length float32, width float32) float32 {
	return length * width
}

func calculateCircleArea(radius float32) float32 {
	const pi float32 = 3.1416

	return pi * radius * radius
}

func calculateTrapezoidArea(minorBase float32, majorBase float32, height float32) float32 {
	return (minorBase + majorBase) * height / 2
}

func main() {
	printHelloWorld()

	printMessage("Learning by doing")

	printParams(10, 10, "Metro Exodus")

	result := multiply(9, 7)
	fmt.Println("Result:", result)

	name, age := getNameAndAge()
	fmt.Println("Name:", name, "Age:", age)

	// (_) is used to ignore function return values
	onlyName, _ := getNameAndAge()
	fmt.Println("Only name:", onlyName)

	// Challenge
	// Calculate the area of a rectangle, circle and trapezoid using functions

	// Rectangle
	rectangleArea := calculateRectangleArea(45, 36)
	fmt.Println("Rectangle area:", rectangleArea)

	// Circle
	circleArea := calculateCircleArea(30)
	fmt.Println("Circle area:", circleArea)

	// Trapezoid
	trapezoidArea := calculateTrapezoidArea(29, 33, 14)
	fmt.Println("Trapezoid area:", trapezoidArea)
}
