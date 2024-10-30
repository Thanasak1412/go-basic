package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("Calculator")

	fmt.Println("Which operation do you want to perform?, (add , subtract, multiply, divide)")
	_, _ = fmt.Scanf("%s", &operation)

	fmt.Println("Enter first number")
	_, _ = fmt.Scanf("%f", &number1)

	fmt.Println("Enter second number")
	_, _ = fmt.Scanf("%f", &number2)

	switch operation {
	case "add":
		fmt.Printf("%d + %d = %d\n", number1, number2, number1+number2)
	case "subtract":
		fmt.Printf("%d - %d = %d\n", number1, number2, number1-number2)
	case "multiply":
		fmt.Printf("%d * %d = %d\n", number1, number2, number1*number2)
	case "divide":
		fmt.Printf("%d / %d = %d\n", number1, number2, number1/number2)
	default:
		fmt.Println("Invalid operation")
	}
}
