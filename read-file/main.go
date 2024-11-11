package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	rootPath, _ := os.Getwd()
	rootPath = fmt.Sprintf("%s/read-file", rootPath)
	nameFile := os.Args[1]

	data := make([]byte, 32*1024)
	file, err := os.Open(fmt.Sprintf("%s/%s", rootPath, nameFile))

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	count, err := file.Read(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
