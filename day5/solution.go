package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	updates, rules := parseData()
	middlePageSum := calcMiddlePageNumberSum(updates, rules)
	fmt.Println("Day 5, Solution 1 result: ", middlePageSum)
	fmt.Println("Day 5, Solution 2 result: ")
}

type rule struct {
	before int
	after  int
}

func parseData() ([][]int, []rule) {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		panic("Failed to open file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make([]rule, 0, 5)
	updates := make([][]int, 0, 5)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parsedRule := strings.Split(line, "|")
		before, err1 := strconv.Atoi(parsedRule[0])
		after, err2 := strconv.Atoi(parsedRule[1])
		if err1 != nil || err2 != nil {
			panic("Failed to parse rule!")
		}

		rules = append(rules, rule{before: before, after: after})
	}

	for scanner.Scan() {
		line := scanner.Text()
		update := make([]int, 0, 5)
		parsedUpdate := strings.Split(line, ",")
		for _, val := range parsedUpdate {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic("Failed to parse update!")
			}

			update = append(update, num)
		}

		updates = append(updates, update)
	}

	return updates, rules
}

func calcMiddlePageNumberSum(updates [][]int, rules []rule) int {
	sum := 0
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			mid := len(update) / 2
			sum += update[mid]
		}
	}

	return sum
}

func isValidUpdate(update []int, rules []rule) bool {
	indexMap := make(map[int]int)
	for i, val := range update {
		indexMap[val] = i
	}

	for _, r := range rules {
		_, okBefore := indexMap[r.before]
		_, okAfter := indexMap[r.after]
		if !okBefore || !okAfter {
			continue
		}

		if indexMap[r.before] > indexMap[r.after] {
			return false
		}
	}

	return true
}
