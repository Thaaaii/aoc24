package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solution() {
	data := parseData()
	t := newTopographicMap(data)
	scoreSum := t.calcScoreSum()
	fmt.Println("Day 10, Solution 1 result: ", scoreSum)
	ratingSum := t.calcRatingSum()
	fmt.Println("Day 10, Solution 2 result: ", ratingSum)
}

func parseData() [][]int {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}

		data = append(data, row)
	}

	return data
}

type TopographicMap struct {
	field   [][]int
	m       int
	n       int
	visited map[string]struct{}
}

func newTopographicMap(data [][]int) TopographicMap {
	return TopographicMap{
		field:   data,
		m:       len(data),
		n:       len(data[0]),
		visited: map[string]struct{}{},
	}
}

func (t TopographicMap) calcScoreSum() int {
	sum := 0
	for i, row := range t.field {
		for j, cell := range row {
			if cell == 0 {
				t.visited = make(map[string]struct{})
				sum += t.dfsTrailScore(i, j, 0)
			}
		}
	}

	return sum
}

func (t TopographicMap) dfsTrailScore(i, j, height int) int {
	if i < 0 || i >= t.m || j < 0 || j >= t.n || height != t.field[i][j] {
		return 0
	}

	if height == 9 {
		if _, ok := t.visited[fmt.Sprintf("%d-%d", i, j)]; ok {
			return 0
		}

		t.visited[fmt.Sprintf("%d-%d", i, j)] = struct{}{}

		return 1
	}

	height++

	return t.dfsTrailScore(i+1, j, height) + t.dfsTrailScore(i-1, j, height) + t.dfsTrailScore(i, j+1, height) + t.dfsTrailScore(i, j-1, height)
}

func (t TopographicMap) calcRatingSum() int {
	sum := 0
	for i, row := range t.field {
		for j, cell := range row {
			if cell == 0 {
				sum += t.dfsTrailRating(i, j, 0)
			}
		}
	}

	return sum
}

func (t TopographicMap) dfsTrailRating(i, j, height int) int {
	if i < 0 || i >= t.m || j < 0 || j >= t.n || height != t.field[i][j] {
		return 0
	}

	if height == 9 {
		return 1
	}

	height++

	return t.dfsTrailRating(i+1, j, height) + t.dfsTrailRating(i-1, j, height) + t.dfsTrailRating(i, j+1, height) + t.dfsTrailRating(i, j-1, height)
}
