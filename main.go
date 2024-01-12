package main

import (
	"aoc2023/days"
	"aoc2023/days/day1"
	"aoc2023/days/day2"
	"aoc2023/days/day3"
	"aoc2023/days/day4"
	"aoc2023/days/day5"
	"aoc2023/days/day6"
	"aoc2023/days/day7"
	"fmt"
	"os"
	"strconv"
	"time"
)

var solutions = [25]days.Day{
	day1.Solution,
	day2.Solution,
	day3.Solution,
	day4.Solution,
	day5.Solution,
	day6.Solution,
	day7.Solution,
}

func main() {
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("first argument must be a valid integer"))
	}
	if day > len(solutions) {
		panic(fmt.Errorf("only days 1 -> %d implemented", len(solutions)))
	}

	startTime := time.Now()
	part1 := solutions[day-1].Part1()
	fmt.Printf("Part 1: %s\t Solved in %v\n", part1, time.Now().Sub(startTime))

	startTime = time.Now()
	part2 := solutions[day-1].Part2()
	fmt.Printf("Part 2: %s\t Solved in %v\n", part2, time.Now().Sub(startTime))
}
