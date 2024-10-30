package main

import "fmt"

func main() {
	fmt.Println("Using a Variadic Function")
	fmt.Println("variadic function with default parameter => ", VariadicGreet())
	fmt.Println("variadic function with passing argument => ", VariadicGreet("John"))

	fmt.Println("Using Zero Values")
	fmt.Println("Zero Values with passing a zero value => ", ZeroValues(""))
	fmt.Println("Zero Values with passing a non-zero value => ", ZeroValues("Alice"))

	fmt.Println("Using a Struct for Configuration")
	fmt.Println("Struct for configuration default parameters", StructConfig(Config{}))
	fmt.Println(StructConfig(Config{Name: "Eve", Age: 25}))
	fmt.Println(StructConfig(Config{Name: "Charlie"}))
}
