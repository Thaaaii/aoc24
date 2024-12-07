package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	equations := parseData()
	firstOperatorSet := []string{"+", "*"}
	validEquationSum := calcSumOfValidEquations(firstOperatorSet, equations)
	fmt.Println("Day 7, Solution 1 result: ", validEquationSum)
	secondOperatorSet := []string{"+", "*", "|"}
	calibratedValidEquationSum := calcSumOfValidEquations(secondOperatorSet, equations)
	fmt.Println("Day 7, Solution 2 result: ", calibratedValidEquationSum)
}

func parseData() []equation {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var equations []equation
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		eq := equation{result: result, nums: []int{}}
		nums := strings.Fields(parts[1])
		for _, num := range nums {
			parsedNum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			eq.nums = append(eq.nums, parsedNum)
		}

		equations = append(equations, eq)
	}

	return equations
}

func operatorPermutation(operators []string, length int) []string {
	if length <= 0 {
		return []string{""}
	}

	var results []string
	for _, operator := range operators {
		subPermutations := operatorPermutation(operators, length-1)
		for _, sub := range subPermutations {
			results = append(results, operator+sub)
		}
	}

	return results
}

type equation struct {
	result int
	nums   []int
}

func calcSumOfValidEquations(operators []string, equations []equation) int {
	sum := 0
	for _, eq := range equations {
		reverse := make([]int, 0, len(eq.nums))
		for i := len(eq.nums) - 1; i >= 0; i-- {
			reverse = append(reverse, eq.nums[i])
		}

		permutations := operatorPermutation(operators, len(eq.nums)-1)
		for _, perm := range permutations {
			stack := make([]int, len(reverse))
			copy(stack, reverse)
			for i, _ := range perm {
				a := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res := executeOperation(perm[i], a, b)
				stack = append(stack, res)
			}

			if stack[0] == eq.result {
				sum += eq.result
				break
			}
		}
	}

	return sum
}

func executeOperation(op byte, a, b int) int {
	switch op {
	case '+':
		return a + b
	case '*':
		return a * b
	case '|':
		left := strconv.Itoa(a)
		right := strconv.Itoa(b)
		concat := left + right
		newNum, err := strconv.Atoi(concat)
		if err != nil {
			panic(err)
		}
		return newNum
	default:
		return 0
	}
}
