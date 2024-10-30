package fileutils

import "os"

func ReadTextFile(path string) (string, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(content), nil
}
