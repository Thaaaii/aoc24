package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Solution() {
	data := parseData()
	matrixData := newMatrix(data)
	sumXmas := matrixData.countXmas()
	fmt.Println("Day 4, Solution 1 result: ", sumXmas)
	fmt.Println("Day 4, Solution 2 result: ")
}

func parseData() []string {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows := make([]string, 0, 5)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}

	return rows
}

type matrix struct {
	data []string
	m    int
	n    int
}

func newMatrix(data []string) matrix {
	return matrix{
		data: data,
		m:    len(data),
		n:    len(data[0]),
	}
}

func (m matrix) countXmas() int {
	sum := 0
	for _, row := range m.data {
		for _, char := range row {
			if string(char) == "X" {
				sum += m.checkEveryDirection()
			}
		}
	}

	return sum
}

func (m matrix) checkEveryDirection() int {
	sum := 0

	// diagonal

	// vertical

	// horizontal

	return sum
}
