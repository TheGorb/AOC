package main

import (
	"AOC/utils"
)

const word = "MAS"
const revert_word = "SAM"

func wordsFounds(i, y, rows, cols int, grid [][]rune) bool {
	count := 0

	directions := [][2]int{
		{1, 1},
		{-1, -1},
		{1, 1},
		{1, -1},
	}

	for _, dir := range directions {
		dirX, dirY := dir[0], dir[1]
		currentCandidate := ""

		for k := 0; k < 3; k++ {
			ni := i + k*dirX
			ny := y + k*dirY
			if ni < 0 || ni >= rows || ny < 0 || ny >= cols {
				break
			}
			currentCandidate += string(grid[ni][ny])
		}

		if currentCandidate == word || currentCandidate == revert_word {
			count++
		}
	}

	if count > 0 {
		return true
	}
	return false
}

func diagRight(i, y, rows, cols int, grid [][]rune) bool {
	currentCandidate := ""

	if i+2 >= rows || y+2 >= cols {
		return false
	}

	currentCandidate += string(grid[i][y])
	currentCandidate += string(grid[i+1][y+1])
	currentCandidate += string(grid[i+2][y+2])

	if currentCandidate == word || currentCandidate == revert_word {
		return true
	}
	return false
}

func diagLeft(i, y, cols int, grid [][]rune) bool {
	currentCandidate := ""

	if i+2 > cols || y-2 < 0 {
		return false
	}

	currentCandidate += string(grid[i][y])
	currentCandidate += string(grid[i+1][y-1])
	currentCandidate += string(grid[i+2][y-2])

	if currentCandidate == word || currentCandidate == revert_word {
		return true
	}
	return false
}

func main() {
	fileContent := utils.ReadFile("day4/input2.txt")
	grid, rows, cols := utils.ParseGrid(fileContent)
	count := 0

	for i := 1; i < rows; i++ {
		for y := 1; y < cols; y++ {
			if grid[i][y] == 'A' {
				if diagRight(i-1, y-1, rows, cols, grid) && diagLeft(i-1, y+1, cols, grid) {
					count++
				}
			}
		}
	}
	println(count)
}
