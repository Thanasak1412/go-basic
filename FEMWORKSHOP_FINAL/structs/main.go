package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

// Area Method with value receiver
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Scale method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 10, height: 5}

	fmt.Println("Area: ", rect.Area())

	rect.Scale(2)
	fmt.Println("Scale Width: ", rect.width, "Scale Heigth: ", rect.height)
}
