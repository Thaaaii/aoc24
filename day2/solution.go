package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	reports := parseInput()
	safeReports := calcSafeReports(reports)
	fmt.Println("Day 2, Solution 1 result: ", safeReports)
	dampenedSafeReports := calcDampenedSafeReports(reports)
	fmt.Println("Day 2, Solution 2 result: ", dampenedSafeReports)
}

func parseInput() [][]int {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reports := make([][]int, 0, 5)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		reportLine := strings.Fields(line)
		levels := make([]int, 0, 5)

		for _, element := range reportLine {
			level, err := strconv.Atoi(element)
			if err != nil {
				panic("Error parsing report")
			}
			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}

	return reports
}

func calcSafeReports(reports [][]int) int {
	sum := 0
	for _, report := range reports {
		if isSafe(report) {
			sum++
		}
	}

	return sum
}

func isSafe(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	checkSafety := func(a, b int) bool {
		return a < b && abs(a-b) <= 3
	}

	if report[0] > report[1] {
		checkSafety = func(a, b int) bool {
			return a > b && abs(a-b) <= 3
		}
	}

	for i := 1; i < len(report); i++ {
		if !checkSafety(report[i-1], report[i]) {
			return false
		}
	}

	return true
}

func calcDampenedSafeReports(reports [][]int) int {
	sum := 0
	for _, report := range reports {
		if isSafeDampened(report) {
			sum++
		}
	}

	return sum
}

func isSafeDampened(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	checkSafety := func(a, b int) bool {
		return a < b && abs(a-b) <= 3
	}

	if report[0] > report[1] {
		checkSafety = func(a, b int) bool {
			return a > b && abs(a-b) <= 3
		}
	}

	for i := 1; i < len(report); i++ {
		if !checkSafety(report[i-1], report[i]) {
			safety1 := append(report[:i-1], report[i:]...)
			safety2 := append(report[:i], report[i+1:]...)
			if isSafe(safety1) || isSafe(safety2) {
				fmt.Println(report, safety1, safety2)
				return true
			}

			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
