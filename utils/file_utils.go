package utils

import (
	"log"
	"os"
)

func ReadFile(filepath string) string {
	data, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return string(data)
}
