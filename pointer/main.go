package main

import (
	"fmt"
)

type Counter struct {
	Value int
}

func main() {
	x := 10
	ptr := &x

	fmt.Println("Outputs the memory address of x => ", ptr)
	fmt.Println("Outputs the value of x => ", *ptr)

	*ptr = 20 // Changes the value of x through the pointer
	fmt.Println("Outputs: ", x)

	Increment(ptr)
	fmt.Println("Increment => ", x)

	Increment(&x)
	fmt.Println("Increment => ", x)

	count := Counter{
		Value: 10,
	}

	count.increment(120)
	count.printValue()

	originalSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("OriginalSlice: ", originalSlice)

	//	pass the slice to the function
	modifySlice(originalSlice)

	fmt.Println("After modification: ", originalSlice)

	z := 10

	fmt.Println(*(&(z)))
}
