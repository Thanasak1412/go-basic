package main

import "fmt"

type Config struct {
	Name string
	Age  int
}

func StructConfig(cfg Config) string {
	if cfg.Name == "" {
		cfg.Name = "Guest"
	}

	if cfg.Age == 0 {
		cfg.Age = 30
	}

	return fmt.Sprintf("%s is %d years old", cfg.Name, cfg.Age)
}
