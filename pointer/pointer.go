package main

import "fmt"

func Increment(num *int) {
	*num += 1
}

func (c *Counter) increment(plusValue int) {
	c.Value += plusValue
}

func (c *Counter) printValue() {
	fmt.Println("Value:", c.Value)
}

func modifySlice(s []int) {
	//	Modifying elements of the slice (the underlying array is affected)
	for i := range s {
		s[i] *= 2
	}

	s = append(s, 20)

	fmt.Println(s)
}
