package filemanager

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ReadLines(pathFile string) (string, error) {
	currPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)

		return "", err
	}

	pricesPath := fmt.Sprintf("%s/%s", currPath, pathFile)

	content, err := os.ReadFile(pricesPath)

	if err != nil {
		fmt.Println("Failed to read prices.txt", err)

		return "", err
	}

	return string(content), nil
}

func WriteJson(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(data)

	if err != nil {
		return err
	}

	return nil
}
