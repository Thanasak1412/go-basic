package main

import (
	"fmt"
	"strings"
)

func main() {
	var sort = "id-asc"

	splited := strings.Split(sort, "-")

	fmt.Println(splited[0])
}
