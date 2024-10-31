package main

import "fmt"

type Number []int

func checkNumber() {
	numbers := make(Number, 11)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("%d is even \n", num)
		} else {
			fmt.Printf("%d odd \n", num)
		}
	}
}
