package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Solution() {
	data := parseData()
	matrixData := newMatrix(data)
	sumXMAS := matrixData.countXmas()
	fmt.Println("Day 4, Solution 1 result: ", sumXMAS)
	sumX_MAS := matrixData.countCrossXmas()
	fmt.Println("Day 4, Solution 2 result: ", sumX_MAS)
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
	for i, row := range m.data {
		for j, char := range row {
			if string(char) == "X" {
				sum += m.checkEveryDirectionXmas(i, j)
			}
		}
	}

	return sum
}

func (m matrix) checkEveryDirectionXmas(i, j int) int {
	sum := 0
	wordBuffer := []byte{0, 0, 0, 0}

	// diagonal: up right
	if i-3 >= 0 && j+3 < m.n {
		wordBuffer[0] = m.data[i][j]
		wordBuffer[1] = m.data[i-1][j+1]
		wordBuffer[2] = m.data[i-2][j+2]
		wordBuffer[3] = m.data[i-3][j+3]
		if string(wordBuffer) == "XMAS" {
			sum++
		}
	}

	// diagonal: down right
	if i+3 < m.m && j+3 < m.n {
		wordBuffer[0] = m.data[i][j]
		wordBuffer[1] = m.data[i+1][j+1]
		wordBuffer[2] = m.data[i+2][j+2]
		wordBuffer[3] = m.data[i+3][j+3]
		if string(wordBuffer) == "XMAS" {
			sum++
		}
	}

	// diagonal: down left
	if i+3 < m.m && j-3 >= 0 {
		wordBuffer[0] = m.data[i][j]
		wordBuffer[1] = m.data[i+1][j-1]
		wordBuffer[2] = m.data[i+2][j-2]
		wordBuffer[3] = m.data[i+3][j-3]
		if string(wordBuffer) == "XMAS" {
			sum++
		}
	}

	// diagonal: up left
	if i-3 >= 0 && j-3 >= 0 {
		wordBuffer[0] = m.data[i][j]
		wordBuffer[1] = m.data[i-1][j-1]
		wordBuffer[2] = m.data[i-2][j-2]
		wordBuffer[3] = m.data[i-3][j-3]
		if string(wordBuffer) == "XMAS" {
			sum++
		}
	}

	// vertical: left
	if j-3 >= 0 {
		wordBuffer[0] = m.data[i][j]
		wordBuffer[1] = m.data[i][j-1]
		wordBuffer[2] = m.data[i][j-2]
		wordBuffer[3] = m.data[i][j-3]
		if string(wordBuffer) == "XMAS" {
			sum++
		}
	}

	// vertical: right
	if j+3 < m.n {
		wordBuffer[0] = m.data[i][j]
		wordBuffer[1] = m.data[i][j+1]
		wordBuffer[2] = m.data[i][j+2]
		wordBuffer[3] = m.data[i][j+3]
		if string(wordBuffer) == "XMAS" {
			sum++
		}
	}

	// horizontal: up
	if i-3 >= 0 {
		wordBuffer[0] = m.data[i][j]
		wordBuffer[1] = m.data[i-1][j]
		wordBuffer[2] = m.data[i-2][j]
		wordBuffer[3] = m.data[i-3][j]
		if string(wordBuffer) == "XMAS" {
			sum++
		}
	}

	// horizontal: down
	if i+3 < m.m {
		wordBuffer[0] = m.data[i][j]
		wordBuffer[1] = m.data[i+1][j]
		wordBuffer[2] = m.data[i+2][j]
		wordBuffer[3] = m.data[i+3][j]
		if string(wordBuffer) == "XMAS" {
			sum++
		}
	}

	return sum
}

func (m matrix) countCrossXmas() int {
	sum := 0
	for i, row := range m.data {
		for j, char := range row {
			if string(char) == "A" {
				sum += m.checkEveryDirectionCrossXmas(i, j)
			}
		}
	}

	return sum
}

func (m matrix) checkEveryDirectionCrossXmas(i, j int) int {
	sum := 0
	wordBuffer := []byte{0, 0, 0}

	if i-1 >= 0 && j-1 >= 0 && i+1 < m.m && j+1 < m.n {
		// diagonal: right up-left down
		wordBuffer[0] = m.data[i-1][j+1]
		wordBuffer[1] = m.data[i][j]
		wordBuffer[2] = m.data[i+1][j-1]
		if string(wordBuffer) == "MAS" || string(wordBuffer) == "SAM" {
			sum++
		}

		// diagonal: left up-right down
		wordBuffer[0] = m.data[i-1][j-1]
		wordBuffer[1] = m.data[i][j]
		wordBuffer[2] = m.data[i+1][j+1]
		if string(wordBuffer) == "MAS" || string(wordBuffer) == "SAM" {
			sum++
		}
	}

	if sum == 2 {
		return 1
	}

	return 0
}
