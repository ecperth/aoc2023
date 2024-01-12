package day1

import (
	"aoc2023/days"
	"aoc2023/utils"
	"strconv"
	"strings"
	"unicode"
)

var input = utils.ReadInputAsStrings(1)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {
	sum := 0
	for _, line := range input {
		firstInt, lastInt := 'a', 'a'
		for _, value := range line {
			if unicode.IsDigit(value) {
				lastInt = value
				if firstInt == 'a' {
					firstInt = value
				}
			}
		}
		lineValue, _ := strconv.Atoi(string(firstInt) + string(lastInt))
		sum += lineValue
	}
	return strconv.Itoa(sum)
}

func part2() string {
	numbers := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	sum := 0
	for _, line := range input {
		firstInt, lastInt := 'a', 'a'
		numberBuffer := ""
		for _, value := range line {
			if unicode.IsDigit(value) {
				numberBuffer = ""
				lastInt = value
				if firstInt == 'a' {
					firstInt = value
				}
			} else {
				numberBuffer += string(value)
				anyCandidate := false
				for k, v := range numbers {
					if strings.HasPrefix(k, numberBuffer) {
						if numberBuffer == k {
							lastInt = v
							if firstInt == 'a' {
								firstInt = v
							}

							numberBuffer = numberBuffer[1:]
							for !anyCandidate {
								numberBuffer = numberBuffer[1:]
								for k, _ := range numbers {
									if strings.HasPrefix(k, numberBuffer) {
										anyCandidate = true
										break
									}
								}
							}
						}
						anyCandidate = true
						break
					}
				}
				if !anyCandidate {
					for !anyCandidate {
						numberBuffer = numberBuffer[1:]
						for k, _ := range numbers {
							if strings.HasPrefix(k, numberBuffer) {
								anyCandidate = true
								break
							}
						}
					}
				}
			}
		}
		lineValue, _ := strconv.Atoi(string(firstInt) + string(lastInt))
		sum += lineValue
	}
	return strconv.Itoa(sum)
}
