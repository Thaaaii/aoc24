package day13

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solution() {
	machines := parseData()
	fmt.Println("Day 13, Solution 1 result: ", calcMinimalTokenUsage(machines, false))
	fmt.Println("Day 13, Solution 2 result: ", calcMinimalTokenUsage(machines, true))
}

type machine struct {
	ax     int
	ay     int
	bx     int
	by     int
	prizeX int
	prizeY int
}

func parseData() []machine {
	file, err := os.Open("day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var machines []machine
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	reg := regexp.MustCompile(`\d+`)
	machineDescription := strings.Split(string(content), "\n\n")

	for _, m := range machineDescription {
		values := reg.FindAllString(m, -1)
		nums := make([]int, 0, 6)
		for _, numString := range values {
			num, err := strconv.Atoi(numString)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}

		machines = append(machines, machine{
			ax:     nums[0],
			ay:     nums[1],
			bx:     nums[2],
			by:     nums[3],
			prizeX: nums[4],
			prizeY: nums[5],
		})

	}

	return machines
}

func calcMinimalTokenUsage(machines []machine, correction bool) int {
	sum := 0
	for _, m := range machines {
		det := m.ax*m.by - m.bx*m.ay
		if det == 0 {
			continue
		}

		if correction {
			m.prizeX += 10000000000000
			m.prizeY += 10000000000000
		}

		x1 := m.prizeX*m.by - m.prizeY*m.bx
		x2 := m.ax*m.prizeY - m.ay*m.prizeX
		a := x1 / det
		b := x2 / det

		if m.ax*a+m.bx*b == m.prizeX && m.ay*a+m.by*b == m.prizeY {
			sum += 3*a + b
		}
	}

	return sum
}
