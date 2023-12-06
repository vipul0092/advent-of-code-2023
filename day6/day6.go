package day6

import (
	"advent-of-code/reader"
	"fmt"
	"strconv"
	"strings"
)

var input = `Time:      7  15   30
Distance:  9  40  200`

func Solve() {
	input = reader.Read("/day6/input")
	lines := strings.Split(input, "\n")

	times := getNumbers(lines[0])
	distances := getNumbers(lines[1])

	var part1 int64 = 1
	for i := 0; i < len(times); i++ {
		part1 *= getWays(int64(times[i]), int64(distances[i]))
	}

	time := getNumber(lines[0])
	distance := getNumber(lines[1])
	part2 := getWays(time, distance)

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}

func getNumbers(line string) []int {
	numbers := make([]int, 0)
	nstr := strings.Split(strings.Split(line, ":")[1], " ")

	for _, ns := range nstr {
		if len(ns) != 0 {
			num, _ := strconv.Atoi(ns)
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func getNumber(line string) int64 {
	var sb strings.Builder
	for _, char := range line {
		if char >= '0' && char <= '9' {
			sb.WriteRune(char)
		}
	}
	num, _ := strconv.ParseInt(sb.String(), 10, 64)
	return num
}

// Binary search on answer
func getWays(time, distance int64) int64 {
	var mini int64 = 1
	maxi := time / 2

	var t1 int64 = 1
	for mini <= maxi {
		mid := (mini + maxi) >> 1
		dist := (time - mid) * mid
		if dist > distance {
			maxi = mid - 1
			t1 = mid
		} else {
			mini = mid + 1
		}
	}

	var t2 int64 = 1
	mini = time / 2
	maxi = time
	for mini <= maxi {
		mid := (mini + maxi) >> 1
		dist := (time - mid) * mid
		if dist > distance {
			mini = mid + 1
			t2 = mid
		} else {
			maxi = mid - 1
		}
	}
	return t2 - t1 + 1
}
