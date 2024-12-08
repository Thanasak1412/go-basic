package main

import (
	"fmt"
	"reflect"
)

func main() {
	startingNumber := 10
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("sum => ", sumUp(startingNumber, numbers...))
	fmt.Println("sum => ", sumUp(startingNumber+10, 10, 20, 40, 50, 100, 200))
}

func sumUp(startingNumber int, numbers ...int) int {
	fmt.Println("typeof numbers => ", reflect.TypeOf(numbers))

	sum := startingNumber

	for _, number := range numbers {
		sum += number
	}

	return sum
}
