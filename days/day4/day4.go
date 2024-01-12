package day4

import (
	"aoc2023/days"
	"aoc2023/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(4)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {
	score := 0
	var myNumbers, winningNumbers []string
	for _, card := range input {
		numbers := strings.Split(card, ": ")[1]
		splitNumbers := strings.Split(numbers, " | ")
		winningNumbers = utils.StringSplitConsecutive(splitNumbers[0], ' ')
		myNumbers = utils.StringSplitConsecutive(splitNumbers[1], ' ')
		winners := 0
		for _, winningNumber := range winningNumbers {
			if slices.Contains(myNumbers, winningNumber) {
				winners++
			}
		}
		if winners > 0 {
			score += int(math.Pow(2, float64(winners-1)))
		}
	}
	return strconv.Itoa(score)
}

func part2() string {
	cardCounts := make(map[int]int)
	var myNumbers, winningNumbers []string
	for cardNumber, card := range input {
		cardCounts[cardNumber] = cardCounts[cardNumber] + 1
		numbers := strings.Split(card, ": ")[1]
		splitNumbers := strings.Split(numbers, " | ")
		winningNumbers = utils.StringSplitConsecutive(splitNumbers[0], ' ')
		myNumbers = utils.StringSplitConsecutive(splitNumbers[1], ' ')
		winners := 0
		for _, winningNumber := range winningNumbers {
			if slices.Contains(myNumbers, winningNumber) {
				winners++
			}
		}

		if winners > 0 {
			for i := 1; i <= winners; i++ {
				cardCounts[cardNumber+i] = cardCounts[cardNumber+i] + cardCounts[cardNumber]
			}
		}
	}

	cardCount := 0
	for cardNumber := 0; cardNumber < len(input); cardNumber++ {
		cardCount = cardCount + cardCounts[cardNumber]
	}
	return strconv.Itoa(cardCount)
}
