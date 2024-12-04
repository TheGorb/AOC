package utils

import (
	"log"
	"os"
	"strings"
)

func ReadFile(filepath string) string {
	data, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return string(data)
}

func GetDimensions(fileContent string) (rows int, cols int) {
	lines := strings.Split(strings.TrimSpace(fileContent), "\n")
	rows = len(lines)
	if rows > 0 {
		cols = len([]rune(lines[0]))
	}
	return rows, cols
}

func ParseGrid(fileContent string) (grid [][]rune, rows int, cols int) {
	lines := strings.Split(strings.TrimSpace(fileContent), "\n")
	rows = len(lines)
	if rows == 0 {
		return nil, 0, 0
	}

	grid = make([][]rune, rows)
	cols = len([]rune(lines[0])) - 1

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid, rows, cols
}
