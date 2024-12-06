package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Solution() {
	data := parseData()
	initialX, initialY := getStartingPosition(data)
	firstMapResult, uniqueCellsVisited := simulateGuardPath(data)
	fmt.Println("Day 6, Solution 1 result: ", uniqueCellsVisited)
	spots := getAllVisitedSpots(initialX, initialY, firstMapResult)
	loopSpots := calcAllLoopSpots(data, spots)
	fmt.Println("Day 6, Solution 2 result: ", loopSpots)
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
		buf := make([]byte, len(row))
		copy(buf, row)
		data = append(data, buf)
	}

	return data
}

func cloneGrid(grid [][]byte) [][]byte {
	var out [][]byte
	for _, row := range grid {
		clone := make([]byte, len(row))
		copy(clone, row)
		out = append(out, clone)
	}

	return out
}

type simulator struct {
	posX      int
	posY      int
	dimX      int
	dimY      int
	grid      [][]byte
	direction byte
}

func newSimulator(data [][]byte) simulator {
	posX, posY := getStartingPosition(data)
	return simulator{
		posX:      posX,
		posY:      posY,
		dimX:      len(data[0]),
		dimY:      len(data),
		grid:      data,
		direction: data[posY][posX],
	}
}

func getStartingPosition(data [][]byte) (int, int) {
	for y, row := range data {
		for x, cell := range row {
			if cell == '^' {
				return x, y
			}
		}
	}
	return 0, 0
}

func simulateGuardPath(data [][]byte) ([][]byte, int) {
	grid := cloneGrid(data)
	s := newSimulator(grid)
	sum := 1
	for {
		dirX, dirY := s.moveIntoDirection()
		newX, newY := s.posX+dirX, s.posY+dirY
		if newY < 0 || newX < 0 || newX == s.dimX || newY == s.dimY {
			s.grid[s.posY][s.posX] = 'X'
			break
		}

		newCell := s.grid[newY][newX]
		switch newCell {
		case '#':
			s.direction = s.rotateGuard()
			s.grid[s.posY][s.posX] = s.direction
			continue
		default:
			if newCell == '.' {
				sum++
			}
			s.grid[s.posY][s.posX] = 'X'
			s.posX = newX
			s.posY = newY
			s.grid[newY][newX] = s.direction
		}
	}

	return s.grid, sum
}

// moveIntoDirection returns the pair x, y
func (s *simulator) moveIntoDirection() (int, int) {
	switch s.direction {
	case '^':
		return 0, -1
	case 'v':
		return 0, 1
	case '<':
		return -1, 0
	case '>':
		return 1, 0
	}

	return 0, 0
}

func (s *simulator) rotateGuard() byte {
	switch s.direction {
	case '^':
		return '>'
	case 'v':
		return '<'
	case '<':
		return '^'
	case '>':
		return 'v'
	}

	return 0
}

func getAllVisitedSpots(initialX, initialY int, grid [][]byte) [][]int {
	var result [][]int
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'X' && (initialX != x || initialY != y) {
				result = append(result, []int{x, y})
			}
		}
	}

	return result
}

func calcAllLoopSpots(data [][]byte, spots [][]int) int {
	sum := 0
	for _, spot := range spots {
		grid := cloneGrid(data)
		s := newSimulator(grid)
		x := spot[0]
		y := spot[1]
		s.grid[y][x] = '#'
		if s.hasLoop() {
			sum++
		}
	}

	return sum
}

func (s *simulator) hasLoop() bool {
	visitMap := make(map[string]struct{})
	for {
		stateKey := getStateKey(s.direction, s.posX, s.posY)
		if _, ok := visitMap[stateKey]; ok {
			return true
		}
		visitMap[stateKey] = struct{}{}

		dirX, dirY := s.moveIntoDirection()
		newX, newY := s.posX+dirX, s.posY+dirY
		if newY < 0 || newX < 0 || newX == s.dimX || newY == s.dimY {
			break
		}

		newCell := s.grid[newY][newX]
		switch newCell {
		case '#':
			s.direction = s.rotateGuard()
			s.grid[s.posY][s.posX] = s.direction
			continue
		default:
			s.grid[s.posY][s.posX] = 'X'
			s.posX = newX
			s.posY = newY
			s.grid[newY][newX] = s.direction
		}
	}

	return false
}

func getStateKey(direction byte, x, y int) string {
	return fmt.Sprintf("%d-%d-%d", direction, x, y)
}
