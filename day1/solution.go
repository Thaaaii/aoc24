package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solution() {
	lists := parseInput()

	if len(lists) != 2 {
		panic("Input has not two lists!")
	}

	if len(lists[0]) != len(lists[1]) {
		panic("The two lists do not have the same amount of entries!")
	}

	differences := calcDifferences(lists)
	fmt.Println("Day 1, Solution 1 result: ", differences)
	similarity := calcSimilarity(lists)
	fmt.Println("Day 1, Solution 2 result: ", similarity)
}

func parseInput() [][]int {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pairs := make([][]int, 2)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) != 2 {
			continue
		}

		a, err1 := strconv.Atoi(fields[0])
		b, err2 := strconv.Atoi(fields[1])

		if err1 != nil || err2 != nil {
			continue
		}

		pairs[0] = append(pairs[0], a)
		pairs[1] = append(pairs[1], b)
	}

	return pairs
}

func calcDifferences(lists [][]int) int {
	sort.Ints(lists[0])
	sort.Ints(lists[1])
	sum := 0

	for i := 0; i < len(lists[0]); i++ {
		maximum := max(lists[0][i], lists[1][i])
		minimum := min(lists[0][i], lists[1][i])
		sum += maximum - minimum
	}

	return sum
}

func calcSimilarity(lists [][]int) int {
	countMap := make(map[int]int)

	for _, val := range lists[1] {
		countMap[val]++
	}

	sum := 0
	for _, val := range lists[0] {
		if occurence, ok := countMap[val]; ok {
			sum += occurence * val
		}
	}

	return sum
}
