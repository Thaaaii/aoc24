package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	nums := parseData()
	numberOfStones25 := calcStonesAfterBlinks(25, nums)
	fmt.Println("Day 11, Solution 1 result: ", numberOfStones25)
	numberOfStones75 := calcStonesAfterBlinks(75, nums)
	fmt.Println("Day 11, Solution 2 result: ", numberOfStones75)
}

func parseData() []int {
	file, err := os.Open("day11/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	numbers := strings.Fields(line)
	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	return nums
}

func calcStonesAfterBlinks(n int, nums []int) int {
	sum := 0
	cache := make(map[string]int)
	for _, num := range nums {
		sum += bottomUp(cache, num, n)
	}

	return sum
}

func bottomUp(cache map[string]int, stone, blinks int) int {
	if blinks == 0 {
		return 1
	}

	key := fmt.Sprintf("%d-%d", stone, blinks)
	if v, ok := cache[key]; ok {
		return v
	}

	val := 0
	if stone == 0 {
		val = bottomUp(cache, 1, blinks-1)
	} else if numString := strconv.Itoa(stone); len(numString)%2 == 0 {
		mid := len(numString) / 2
		left, err1 := strconv.Atoi(numString[:mid])
		right, err2 := strconv.Atoi(numString[mid:])
		if err1 != nil || err2 != nil {
			panic("error while converting string to int")
		}
		val = bottomUp(cache, left, blinks-1) + bottomUp(cache, right, blinks-1)
	} else {
		val = bottomUp(cache, stone*2024, blinks-1)
	}

	cache[key] = val
	return val
}
