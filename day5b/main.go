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

// Perform topological sort using re-adapted Kahn's algorithm
// https://www.geeksforgeeks.org/topological-sorting-indegree-based-solution
func reorderUpdate(update []int, rules [][]string) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, page := range update {
		graph[page] = []int{}
		inDegree[page] = 0
	}

	for _, rule := range rules {
		x, _ := strconv.Atoi(rule[0])
		y, _ := strconv.Atoi(rule[1])

		if contains(update, x) && contains(update, y) {
			graph[x] = append(graph[x], y)
			inDegree[y]++
		}
	}

	var queue []int
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	fileContent := utils.ReadFile("day5b/input.txt")
	rulesRaw, updatesRaw := parseInput(fileContent)
	var correctedMiddlePages []int

	for _, update := range updatesRaw {
		if !isValidUpdate(update, rulesRaw) {
			corrected := reorderUpdate(update, rulesRaw)
			middleIndex := len(corrected) / 2
			correctedMiddlePages = append(correctedMiddlePages, corrected[middleIndex])
		}
	}

	sum := 0
	for _, middlePage := range correctedMiddlePages {
		sum += middlePage
	}

	println(sum)
}
