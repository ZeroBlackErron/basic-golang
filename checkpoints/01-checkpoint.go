package main

import "fmt"

func printSomething() {
	fmt.Println("Some random text")
}

func printAMessage(message string) {
	fmt.Println("Some text:", message)
}

func formatMessage(message string) string {
	return fmt.Sprintf("Some text:" + message)
}

func returnMessageWithCode() (string, int) {
	return "An error has occurred", 400
}

func deferExample() {
	defer printSomething()

	printAMessage("In first place, ...")
	printAMessage("Then, ...")
	printAMessage("Finally, ...")
}

func main() {
	fmt.Println("Hello world")

	const gender = "male"
	var name = "Anthony"
	lastName := "Rosado"
	var age uint = 23
	var height float32 = 175
	var profileComplete bool

	fmt.Println("Full name:", name, lastName)
	fmt.Println("Age:", age)
	fmt.Println("Gender:", gender)
	fmt.Println("Height:", height)
	fmt.Println("Single:", profileComplete)

	x := 30
	y := 40

	fmt.Println("x + y:", x+y)
	fmt.Println("x - y:", x-y)
	fmt.Println("x * y:", x*y)
	fmt.Println("x / y:", x/y)
	fmt.Println("x % y:", x%y)
	x++
	fmt.Println("x++:", x)
	y--
	fmt.Println("y--:", y)

	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x <= y:", x <= y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x >= y:", x >= y)

	fmt.Println("x < y || x >= y:", x < y || x >= y)
	fmt.Println("x < y && x >= y:", x < y && x >= y)

	msg := fmt.Sprintf("Full name: %s %s\nAge: %d\nHeight: %v", name, lastName, age, height)
	fmt.Println(msg)

	fmt.Printf("%d < %d || %d >= %d: ", x, y, x, y)
	fmt.Println(x < y || x >= y)

	fmt.Printf("%d < %d && %d >= %d: ", x, y, x, y)
	fmt.Println(x < y && x >= y)

	printSomething()
	printAMessage("Really cool quote")
	formattedMessage := formatMessage("Another message")
	printAMessage(formattedMessage)
	errorMessage, errorCode := returnMessageWithCode()

	fmt.Printf("Message: %s | Code: %d\n", errorMessage, errorCode)

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %d | ", i)
	}

	randomQuote := "random quote"

	for i, l := range randomQuote {
		fmt.Printf("i: %v | l: %v\n", i, l)
	}

	backCounter := 3

	for backCounter > 0 {
		fmt.Printf("count: %d\n", backCounter)
		backCounter--
	}

	if 4 == 5 || 6 > 1 {
		fmt.Println("four is not equal to five and six is greater then one")
	}

	if 3 == 3 || 5 > 6 {
		fmt.Println("three is equal to three and five is not greater then six")
	} else if 4 == 5 || 6 > 1 {
		fmt.Println("four is not equal to five and six is greater then one")
	}

	if 2 == 3 || 5 > 6 {
		fmt.Println("three is equal to three and five is not greater then six")
	} else if 4 == 5 || 6 > 7 {
		fmt.Println("four is not equal to five and six is greater then one")
	} else {
		fmt.Println("else case")
	}

	score := 100

	switch {
	case score > 90:
		fmt.Println("A")
		break
	case score > 60:
		fmt.Println("B")
		break
	case score > 30:
		fmt.Println("C")
		break
	default:
		fmt.Println("D")
		break
	}

	switch score {
	case 50:
		fmt.Println("50/50")
		break
	case 100:
		fmt.Println("Perfect")
		break
	}

	deferExample()

}
