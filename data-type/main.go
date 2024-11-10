package main

import "fmt"

type Dog struct{}

type Animal interface {
	Speak() string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	// Data type
	fmt.Println("Basic Types ===================================>")
	var integer int = 10
	fmt.Println("Integer => use for generate-purpose integer arithmetic", integer)

	var integer8 int8 = 127
	fmt.Println("Integer8 => when need to control the size (-127 through 128)", integer8)

	var integer16 int16 = 32767
	fmt.Println("Integer16 => when need to control the size between -32768 through 32767", integer16)

	var integer32 int32 = -2147483648
	fmt.Println("Integer32 => when need to control the size between -2147483648 through 2147483647", integer32)

	var integer64 int64 = -9223372036854775808
	fmt.Println("Integer64 => when need to control the size between -9223372036854775808 through 9223372036854775807", integer64)

	var f64 float64 = -12392.20
	fmt.Println("Float64 => for most floating-point arithmetic since it provides higher precision", f64)

	var f32 float32 = 3.14159265358979323846
	fmt.Println("Float32 => when working with large datasets that need lower memory usage (graphics programming)", f32)

	var c64 complex64 = 1 + 2i
	var c128 complex128 = 1 + 2i
	fmt.Printf("use these then you need to represent and perform operations on complex numbers (in signal processing, physics simulations) complex64: %f, complex128: %f\n", c64, c128)

	var b byte = 255 // Maximum value for a byte
	fmt.Println("used to represent raw binary data", b)
	//	Byte array (used commonly in Go for raw data
	data := []byte{72, 101, 108, 108, 111}
	fmt.Println("data: ", data)
	fmt.Println("string(data): ", string(data)) // Convert byte slice to string

	var r rune = 'A'
	fmt.Println("r: ", r)
	fmt.Printf("Unicode for 'A': U+%04X\n", r)

	//	Another example with a non-ASCII character
	var emoji rune = '🙂'
	fmt.Println("emoji: ", emoji)
	fmt.Printf("Unicode for '🙂': U+%04X\n", emoji)

	var boolean bool = false
	fmt.Println("boolean: ", boolean)

	var str string = "use for data that won't change (like configuration, logs, or text processing)"
	fmt.Println("string: ", str)

	fmt.Println("Composite Types ======================>")

	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Use arrays when you need a fixed-size collection of elements with the same type: ", arr)

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Slices are used for dynamically-sized and are the most common way to work with collections of elements", slice)

	type Person struct {
		Name string
		Age  int
	}
	var p Person = Person{Name: "Alice", Age: 30}
	fmt.Println("Use structs to group related data in a single type: ", p)
	fmt.Println("Access the value in struct: Name", p.Name)
	fmt.Println("Access the value in struct: Age", p.Age)

	var money float64 = 10.20
	pointer := &money
	*pointer = 200.9999
	fmt.Println("Use pointers when you need to share and manipulate data across different parts of your program without making copies", money)
	fmt.Println("Pointer: Stores the memory address of a value", pointer)

	fmt.Println("Reference Types ==========================>")

	var m = map[string]int{"apple": 5, "oranges": 3}

	fmt.Println("Use maps when you need to store key-value pairs", m)
	fmt.Println("Retrieve value by key", m["apple"])
	fmt.Println("Retrieve value by key", m["oranges"])

	var square func(int) int = func(x int) int { return x * x }
	result := square(10)
	fmt.Println("Functions are first-class citizens in Go", result)

	ch := make(chan int)
	go func() {
		ch <- 1 // Send value into channel
	}()
	value := <-ch // Receive value from channel
	fmt.Println("Use channels for communication between goroutines", value)

	fmt.Println("Interface Types")

	var a Animal = Dog{}   // Dog implements the Animal interface
	fmt.Println(a.Speak()) // Output: Woof!

	animal := make(map[string]int)

	animal["dog"] = 10
	animal["cat"] = 1

	for breed, amount := range animal {
		fmt.Printf("%s has %d\n", breed, amount)
	}

	addAnimal(animal)

	fmt.Println("rat: ", animal["rat"])
}

func addAnimal(a map[string]int) {
	a["rat"] = 2
}
