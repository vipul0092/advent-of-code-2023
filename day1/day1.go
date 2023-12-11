package day1

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

const input = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

var nos = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func Solve() {
	//lines := strings.Split(input, "\n")
	lines := strings.Split(utils.Read("/day1/input"), "\n")

	sum := 0
	for _, line := range lines {
		first := 0
		second := 0
		for _, char := range line {
			if char >= '0' && char <= '9' {
				if first == 0 {
					first = int(char - '0')
				}
				second = int(char - '0')
			}
		}
		val := ((first) * 10) + second
		sum += val
	}

	fmt.Println("Part 1: ", sum)

	sum2 := 0
	for _, line := range lines {
		fidx := -1
		first := 0
		sidx := -1
		second := 0
		for idx, char := range line {
			if char >= '0' && char <= '9' {
				if first == 0 {
					first = int(char - '0')
					fidx = idx
				}
				second = int(char - '0')
				sidx = idx
			}
		}

		for idx, str := range nos {
			fstr := strings.Index(line, str)
			lstr := strings.LastIndex(line, str)

			if fidx == -1 || (fstr != -1 && fstr < fidx) {
				fidx = fstr
				first = idx + 1
			}

			if sidx == -1 || (lstr != -1 && lstr > sidx) {
				sidx = lstr
				second = idx + 1
			}
		}

		val := ((first) * 10) + second
		sum2 += val
	}
	fmt.Println("Part 2: ", sum2)
}
