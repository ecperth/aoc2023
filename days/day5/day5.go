package day5

import (
	"aoc2023/days"
	"aoc2023/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(5)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {
	//Parse
	var seeds []int
	seeds = utils.StringSliceToIntSlice(utils.StringSplitConsecutive(strings.Split(input[0], ": ")[1], ' '))
	var maps [][][]int
	var nextMap [][]int
	for lineNum := 3; lineNum < len(input); lineNum++ {
		if len(input[lineNum]) == 0 {
			maps = append(maps, nextMap)
			nextMap = make([][]int, 0)
			lineNum++
			continue
		}
		nextMap = append(nextMap, utils.StringSliceToIntSlice(utils.StringSplitConsecutive(input[lineNum], ' ')))
	}
	maps = append(maps, nextMap)

	//Solve
	answer := -1
	for _, seed := range seeds {
		value := seed
		for _, step := range maps {
			for _, mapping := range step {
				if value >= mapping[1] && value < mapping[1]+mapping[2] {
					value += (mapping[0] - mapping[1])
					break
				}
			}
		}
		if answer == -1 || value < answer {
			answer = value

		}
	}
	return strconv.Itoa(answer)
}

func part2() string {
	return strconv.Itoa(2)
}
