package day7

import (
	"aoc2023/days"
	"aoc2023/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(7)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {
	time := utils.StringSliceToIntSlice(utils.StringSplitConsecutive(strings.Split(input[0], ":")[1], ' '))
	dist := utils.StringSliceToIntSlice(utils.StringSplitConsecutive(strings.Split(input[1], ":")[1], ' '))

	result := 1
	for i := 0; i < len(time); i++ {
		options := 0
		for t := 1; t < time[i]; t++ {
			if t*(time[i]-t) > dist[i] {
				options++
			}
		}
		result *= options
	}
	return strconv.Itoa(result)
}

func part2() string {
	return strconv.Itoa(2)
}
