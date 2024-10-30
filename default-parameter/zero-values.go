package main

func ZeroValues(name string) string {
	if name == "" {
		name = "Guest"
	}

	return name
}
