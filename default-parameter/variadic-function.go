package main

// VariadicGreet Using a variadic function
// You can use a variadic function to pass on optional argument and set a default value if none is provided.
func VariadicGreet(name ...string) string {
	actualName := "Guest" // Default value

	if len(name) > 0 {
		actualName = name[0]
	}

	return actualName
}
