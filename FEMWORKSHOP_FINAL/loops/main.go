package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "bat")
	animals = append(animals, "elephant")

	animals = slices.Delete(animals, 2, 3)

	for index, value := range animals {
		fmt.Printf("this is my index %d and this is my animal %s\n\n", index, value)
	}

	i := 0
	for i < 5 {
		fmt.Printf("this is a while loop %d\n", i)
		i++
	}
}
