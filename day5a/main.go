package main

import (
	"AOC/utils"
	"strconv"
	"strings"
)

func parseInput(input string) ([][]string, [][]int) {
	sections := strings.Split(input, "\r\n\r\n")
	rulesRaw := strings.Split(sections[0], "\r\n")
	updatesRaw := strings.Split(sections[1], "\r\n")

	var rules [][]string
	for _, rule := range rulesRaw {
		rules = append(rules, strings.Split(rule, "|"))
	}

	var updates [][]int
	for _, update := range updatesRaw {
		parts := strings.Split(update, ",")
		var updateNums []int

		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			updateNums = append(updateNums, num)
		}

		updates = append(updates, updateNums)
	}

	return rules, updates
}

func isValidUpdate(update []int, rules [][]string) bool {
	position := make(map[int]int)

	for i, page := range update {
		position[page] = i
	}

	for _, rule := range rules {
		x, _ := strconv.Atoi(rule[0])
		y, _ := strconv.Atoi(rule[1])

		if posX, existsX := position[x]; existsX {
			if posY, existsY := position[y]; existsY {
				if posX >= posY {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	fileContent := utils.ReadFile("day5a/input.txt")
	rulesRaw, updatesRaw := parseInput(fileContent)

	var validMiddlePages []int

	for _, update := range updatesRaw {
		if isValidUpdate(update, rulesRaw) {
			middleIndex := len(update) / 2
			validMiddlePages = append(validMiddlePages, update[middleIndex])
		}
	}

	sum := 0
	for _, middlePage := range validMiddlePages {
		sum += middlePage
	}

	println(sum)
}
