package main

import (
	"flag"
	"github.com/Thaaaii/aoc24/day1"
	"github.com/Thaaaii/aoc24/day2"
	"github.com/Thaaaii/aoc24/day3"
	"github.com/Thaaaii/aoc24/day4"
)

var (
	solutionsMap = map[int]func(){
		1: day1.Solution,
		2: day2.Solution,
		3: day3.Solution,
		4: day4.Solution,
	}
)

func main() {
	day := flag.Int("day", 1, "prints the solution of the riddle for the given day")
	flag.Parse()
	solutionsMap[*day]()
}