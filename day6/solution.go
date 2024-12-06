package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Solution() {
	data := parseData()
	for i, row := range data {
		if i == 0 || i == 95 || i == 96 {
			fmt.Println(string(row))
		}
	}
	sim := newSimulator(data)
	fmt.Println(sim.dimX, sim.dimY, sim.posX, sim.posY)
	fmt.Println("Day 6, Solution 1 result: ")
	fmt.Println("Day 6, Solution 2 result: ")
}

func parseData() [][]byte {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]byte
	for scanner.Scan() {
		row := scanner.Bytes()
		data = append(data, row)
	}

	return data
}

type simulator struct {
	posX int
	posY int
	dimX int
	dimY int
	grid [][]byte
}

func newSimulator(data [][]byte) simulator {
	posX, posY := getStartingPosition(data)
	return simulator{
		posX: posX,
		posY: posY,
		dimX: len(data[0]),
		dimY: len(data),
		grid: data,
	}
}

func getStartingPosition(data [][]byte) (int, int) {
	for y, row := range data {
		for x, cell := range row {
			if cell != '.' && cell != '#' {
				fmt.Println(cell)
				return x, y
			}
		}
	}
	return 0, 0
}

func (s *simulator) simulateGuardPath() int {
	sum := 0

	return sum
}
