package day12

import (
	"bufio"
	"fmt"
	"os"
)

func Solution() {
	data := parseData()
	gm := newGardenMap(data)
	fmt.Println("Day 12, Solution 1 result: ", gm.calcTotalPrice())
	fmt.Println("Day 12, Solution 2 result: ")
}

func parseData() [][]byte {
	file, err := os.Open("day12/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buf := scanner.Bytes()
		line := make([]byte, len(buf))
		copy(line, buf)
		data = append(data, line)
	}

	return data
}

type gardenMap struct {
	grid       [][]byte
	visited    [][]bool
	areas      map[string]int
	perimeters map[string]int
	m          int
	n          int
}

func newGardenMap(data [][]byte) gardenMap {
	m := len(data)
	n := len(data[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	return gardenMap{
		grid:       data,
		visited:    visited,
		areas:      make(map[string]int),
		perimeters: make(map[string]int),
		m:          m,
		n:          n,
	}
}

func (g *gardenMap) calcTotalPrice() int {
	sum := 0
	for i, row := range g.grid {
		for j, cell := range row {
			if !g.visited[i][j] {
				key := fmt.Sprintf("%d-%d", i, j)
				g.collectAreaInformation(cell, i, j, key)
			}
		}
	}

	for key := range g.areas {
		sum += g.areas[key] * g.perimeters[key]
	}

	return sum
}

func (g *gardenMap) collectAreaInformation(char byte, i, j int, key string) {
	if i < 0 || i >= g.m || j < 0 || j >= g.n || g.grid[i][j] != char {
		g.perimeters[key]++
		return
	}

	if g.visited[i][j] {
		return
	}

	g.visited[i][j] = true
	g.areas[key]++

	g.collectAreaInformation(char, i-1, j, key)
	g.collectAreaInformation(char, i+1, j, key)
	g.collectAreaInformation(char, i, j-1, key)
	g.collectAreaInformation(char, i, j+1, key)
}
