package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solution() {
	data := parseData()
	ids := createFileSystem(data)
	filesystemChecksum := calcFragmentationChecksum(ids)
	fmt.Println("Day 9, Solution 1 result: ", filesystemChecksum)
	fmt.Println("Day 9, Solution 2 result: ")
}

func parseData() []byte {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Bytes()
}

func createFileSystem(data []byte) []int {
	var ids []int
	idCounter := 0
	for i, char := range data {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			for range num {
				ids = append(ids, idCounter)
			}
			idCounter++
		} else {
			for range num {
				ids = append(ids, -1)
			}
		}
	}

	return ids
}

func calcFragmentationChecksum(ids []int) int {
	moveBlocks(ids)

	return checksum(ids)
}

func moveBlocks(ids []int) {
	left, right := 0, len(ids)-1
	for left < right {
		if ids[left] != -1 {
			left++
			continue
		}

		if ids[right] == -1 {
			right--
			continue
		}

		ids[left], ids[right] = ids[right], ids[left]
		left++
		right--
	}
}

func checksum(ids []int) int {
	sum := 0
	for pos, id := range ids {
		if id == -1 {
			continue
		}

		sum += id * pos
	}

	return sum
}
