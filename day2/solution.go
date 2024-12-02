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
	reallySafeReports := calcReallySafeReports(reports)
	fmt.Println("Day 2, Solution 2 result: ", reallySafeReports)
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
	inc := report[1] > report[0]

	if inc {
		for i := 1; i < len(report); i++ {
			diff := report[i] - report[i-1]
			if !(diff >= 1 && diff <= 3) {
				return false
			}
		}

		return true
	}

	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if !(diff >= 1 && diff <= 3) {
			return false
		}
	}

	return true
}

func calcReallySafeReports(reports [][]int) int {
	sum := 0
	for _, report := range reports {
		if isReallySafe(report) {
			sum++
		}
	}

	return sum
}

func isReallySafe(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		modifiedList := make([]int, len(report))
		copy(modifiedList, report)
		modifiedList = append(modifiedList[:i], modifiedList[i+1:]...)
		if isSafe(modifiedList) {
			return true
		}
	}

	return false
}
