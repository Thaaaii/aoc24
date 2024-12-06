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
	incorrectMiddlePageSum := calcIncorrectlyMiddlePageNumberSum(updates, rules)
	fmt.Println("Day 5, Solution 2 result: ", incorrectMiddlePageSum)
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
		indexMap := createIndexMap(update)
		if _, ok := isValidUpdate(indexMap, rules); ok {
			mid := len(update) / 2
			sum += update[mid]
		}
	}

	return sum
}

func createIndexMap(update []int) map[int]int {
	indexMap := make(map[int]int)
	for i, val := range update {
		indexMap[val] = i
	}

	return indexMap
}

// isValidUpdate returns idx of failed rule and bool of operation
func isValidUpdate(indexMap map[int]int, rules []rule) (int, bool) {
	for idx, r := range rules {
		_, okBefore := indexMap[r.before]
		_, okAfter := indexMap[r.after]
		if !okBefore || !okAfter {
			continue
		}

		if indexMap[r.before] > indexMap[r.after] {
			return idx, false
		}
	}

	return -1, true
}

func calcIncorrectlyMiddlePageNumberSum(updates [][]int, rules []rule) int {
	sum := 0
	for _, update := range updates {
		indexMap := createIndexMap(update)
		idx, ok := isValidUpdate(indexMap, rules)
		if ok {
			continue
		}

		for !ok {
			i, j := indexMap[rules[idx].before], indexMap[rules[idx].after]
			indexMap[rules[idx].before], indexMap[rules[idx].after] = j, i
			update[i], update[j] = update[j], update[i]
			idx, ok = isValidUpdate(indexMap, rules)
		}

		mid := len(update) / 2
		sum += update[mid]
	}

	return sum
}
