package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Solution() {
	data := parseFile()
	mulStatemens := parseMulStatementsOnly(string(data))
	sumAll := sumMultiplication(mulStatemens)
	fmt.Println("Day 3, Solution 1 result: ", sumAll)
	allStatements := parseMulAndCondStatements(string(data))
	sumConditional := sumConditionalMultiplication(allStatements)
	fmt.Println("Day 3, Solution 2 result: ", sumConditional)
}

func parseFile() []byte {
	data, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic("Error reading file!")
	}

	return data
}

func parseMulStatementsOnly(data string) []string {
	reg := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	statements := reg.FindAllString(data, -1)
	return statements
}

func parseMulAndCondStatements(data string) []string {
	reg := regexp.MustCompile(`don't\(\)|do\(\)|mul\(\d{1,3},\d{1,3}\)`)
	statements := reg.FindAllString(data, -1)
	return statements
}

func parseNumbersFromStatement(statement string) (int, int) {
	reg := regexp.MustCompile(`\d{1,3}`)
	nums := reg.FindAllString(statement, -1)
	a, err1 := strconv.Atoi(nums[0])
	b, err2 := strconv.Atoi(nums[1])
	if err1 != nil || err2 != nil {
		panic("Error while parsing numbers from mul statment!")
	}

	return a, b
}

func sumMultiplication(statements []string) int {
	sum := 0
	for _, statement := range statements {
		a, b := parseNumbersFromStatement(statement)
		sum += a * b
	}

	return sum
}

func sumConditionalMultiplication(statements []string) int {
	sum := 0
	executeMul := true
	for _, statement := range statements {
		switch statement {
		case "don't()":
			executeMul = false
		case "do()":
			executeMul = true
		default:
			if executeMul {
				a, b := parseNumbersFromStatement(statement)
				sum += a * b
			}
		}
	}

	return sum
}
