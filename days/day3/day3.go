package day3

import (
	"aoc2023/days"
	"aoc2023/utils"
	"slices"
	"strconv"
	"unicode"
)

var input = utils.ReadInputAsStrings(3)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	numbers := make(map[[2]int]string)
	var symbols [][2]int
	width := len(input[0])
	for rowNum, row := range input {
		numberCol := -1
		numberBuffer := ""
		for colNum := 0; colNum < len(row); colNum++ {
			if unicode.IsDigit(rune(row[colNum])) {
				if numberCol != -1 {
					numberBuffer += string(row[colNum])
				} else {
					numberCol = colNum
					numberBuffer = string(row[colNum])
				}
			} else {
				if numberCol != -1 {
					numbers[[2]int{numberCol, rowNum}] = numberBuffer
					numberCol = -1
					numberBuffer = ""
				}
				if row[colNum] != '.' {
					symbols = append(symbols, [2]int{colNum, rowNum})
				}
			}
		}
	}

	sum := 0
	for pos, value := range numbers {
		isPart := false
		minX := max(pos[0]-1, 0)
		maxX := min(pos[0]+len(value), width-1)
		for dy := -1; dy <= 1; dy = dy + 2 {
			y := pos[1] + dy
			if y < 0 || y >= len(input) {
				continue
			}
			for x := minX; x <= maxX; x++ {
				if slices.Contains(symbols, [2]int{x, y}) {
					isPart = true
					break
				}
			}
			if isPart {
				break
			}
		}
		if !isPart {
			if (pos[0] > 0 && slices.Contains(symbols, [2]int{pos[0] - 1, pos[1]})) || (pos[0]+len(value) <= width && slices.Contains(symbols, [2]int{pos[0] + len(value), pos[1]})) {
				isPart = true
			}
		}
		if isPart {
			v, _ := strconv.Atoi(value)
			sum += v
		}
	}
	return strconv.Itoa(sum)
}

func part2() string {
	var empty struct{}

	numbers := make(map[[2]int]string)
	var gearCandidates [][2]int
	width := len(input[0])

	for rowNum, row := range input {
		numberCol := -1
		numberBuffer := ""
		for colNum := 0; colNum < len(row); colNum++ {
			if unicode.IsDigit(rune(row[colNum])) {
				if numberCol != -1 {
					numberBuffer += string(row[colNum])
				} else {
					numberCol = colNum
					numberBuffer = string(row[colNum])
				}
			} else {
				if numberCol != -1 {
					numbers[[2]int{numberCol, rowNum}] = numberBuffer
					numberCol = -1
					numberBuffer = ""
				}
				if row[colNum] == '*' {
					gearCandidates = append(gearCandidates, [2]int{colNum, rowNum})
				}
			}
		}
	}

	sum := 0
	for _, gearCandidate := range gearCandidates {
		adjacentPartNumbers := map[[2]int]struct{}{}
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dy == 0 && dx == 0 {
					continue
				}
				x := gearCandidate[0] + dx
				y := gearCandidate[1] + dy
				if y < 0 || y >= len(input) || x < 0 || x >= width {
					continue
				}
				for pos, value := range numbers {
					isAdjacent := false
					for dx2 := 0; dx2 < len(value); dx2++ {
						if pos[0]+dx2 == x && pos[1] == y {
							isAdjacent = true
							break
						}
					}
					if isAdjacent {
						adjacentPartNumbers[pos] = empty
						break
					}

				}
			}
		}
		if len(adjacentPartNumbers) == 2 {
			ratio := 1
			for k, _ := range adjacentPartNumbers {
				ratio *= utils.AtoiUnsafe(numbers[k])
			}
			sum += ratio
		}
	}
	return strconv.Itoa(sum)
}
