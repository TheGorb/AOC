package main

import (
	"AOC/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const DO = "do()"
const DONT = "don't()"
const MUL = "mul("

func mul(firstNumber int, secondNumber int) int {
	return firstNumber * secondNumber
}

func invalidChar(char rune) bool {
	return !unicode.IsDigit(char) && char != ','
}
func potentialStart(pos int, fileContent string) string {
	var toBeParsed string

	for pos < len(fileContent) && fileContent[pos] != ')' {
		if invalidChar(rune(fileContent[pos])) {
			return ""
		}
		toBeParsed += string(fileContent[pos])
		pos++
	}

	if pos >= len(fileContent) || fileContent[pos] != ')' {
		return ""
	}

	return toBeParsed
}

func parseBothNumber(betweenParenthesis string) (int, int) {
	parts := strings.Split(betweenParenthesis, ",")
	if len(parts) != 2 {
		fmt.Println("Error: Invalid input format, expected two numbers separated by a comma")
		return 0, 0
	}

	firstNumber, err1 := strconv.Atoi(parts[0])
	secondNumber, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid number format")
		return 0, 0
	}

	return firstNumber, secondNumber
}

func main() {
	fileContent := utils.ReadFile("day3/secondInput.txt")
	enabled := true
	total := 0

	for pos := range fileContent {
		if len(fileContent) < pos+5 {
			break
		}

		if matchStringAtPos(pos, fileContent, DO) {
			enabled = true
		}

		if matchStringAtPos(pos, fileContent, DONT) {
			enabled = false
			pos += len(DONT)
		}

		if matchStringAtPos(pos, fileContent, MUL) && enabled {
			pos += len(MUL)
			betweenParenthesis := potentialStart(pos, fileContent)

			if betweenParenthesis != "" {
				firstNumber, secondNumber := parseBothNumber(betweenParenthesis)
				total += mul(firstNumber, secondNumber)
				pos += len(betweenParenthesis)
			}
		}
	}
	fmt.Print(total)
}

func matchStringAtPos(pos int, fileContent string, substr string) bool {
	if pos+len(substr) > len(fileContent) {
		return false
	}
	return fileContent[pos:pos+len(substr)] == substr
}
