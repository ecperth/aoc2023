package day2

import (
	"aoc2023/days"
	"aoc2023/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(2)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	ballCounts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for index, game := range input {
		roundsString := strings.Split(game, ": ")[1]
		rounds := strings.Split(roundsString, "; ")
		possible := true

		for _, round := range rounds {
			balls := strings.Split(round, ", ")
			for _, ball := range balls {
				var count int
				var colour string
				utils.SscanfUnsafe(ball, "%d %s", &count, &colour)
				if ballCounts[colour] < count {
					possible = false
					break
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			sum += index + 1
		}
	}
	return strconv.Itoa(sum)
}

func part2() string {
	sum := 0
	for _, game := range input {
		roundsString := strings.Split(game, ": ")[1]
		rounds := strings.Split(roundsString, "; ")
		minBalls := make(map[string]int)

		for _, round := range rounds {
			balls := strings.Split(round, ", ")
			for _, ball := range balls {
				var count int
				var colour string
				utils.SscanfUnsafe(ball, "%d %s", &count, &colour)
				minBallValue := minBalls[colour]
				if minBallValue < count {
					minBalls[colour] = count
				}
			}
		}
		power := 1
		for _, v := range minBalls {
			power *= v
		}
		sum += power
	}
	return strconv.Itoa(sum)
}
