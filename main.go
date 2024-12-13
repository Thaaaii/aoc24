package main

import (
	"flag"
	"github.com/Thaaaii/aoc24/day12"
	"github.com/Thaaaii/aoc24/day13"

	"github.com/Thaaaii/aoc24/day10"
	"github.com/Thaaaii/aoc24/day11"
	"github.com/Thaaaii/aoc24/day7"
	"github.com/Thaaaii/aoc24/day9"

	"github.com/Thaaaii/aoc24/day1"
	"github.com/Thaaaii/aoc24/day2"
	"github.com/Thaaaii/aoc24/day3"
	"github.com/Thaaaii/aoc24/day4"
	"github.com/Thaaaii/aoc24/day5"
	"github.com/Thaaaii/aoc24/day6"
)

var (
	solutionsMap = map[int]func(){
		1:  day1.Solution,
		2:  day2.Solution,
		3:  day3.Solution,
		4:  day4.Solution,
		5:  day5.Solution,
		6:  day6.Solution,
		7:  day7.Solution,
		9:  day9.Solution,
		10: day10.Solution,
		11: day11.Solution,
		12: day12.Solution,
		13: day13.Solution,
	}
)

func main() {
	day := flag.Int("day", 1, "prints the solution of the riddle for the given day")
	flag.Parse()
	solutionsMap[*day]()
}
