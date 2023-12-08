package day8

import (
	"advent-of-code/reader"
	"fmt"
	"strings"
)

func Solve() {
	lines := reader.ReadLines("/day8/input")
	directions := lines[0]

	paths := make(map[string][]string)
	starts := make([]string, 0)

	for i := 2; i < len(lines); i++ {
		line := lines[i][:len(lines[i])-1]
		parts := strings.Split(line, " = (")
		from := parts[0]
		to := strings.Split(parts[1], ", ")

		paths[from] = to
		if strings.HasSuffix(from, "A") {
			starts = append(starts, from)
		}
	}

	fmt.Println("Part 1: ",
		getCount(func(current string) bool { return current != "ZZZ" }, "AAA", directions, paths))

	counts := make([]int64, len(starts))
	for i, start := range starts {
		count := getCount(func(current string) bool { return !strings.HasSuffix(current, "Z") },
			start, directions, paths)
		counts[i] = int64(count)
	}

	fmt.Println("Part 2: ", lcm(counts))
}

func getCount(loopCheck func(current string) bool, start, directions string, paths map[string][]string) int {
	steps := 0
	dir := 0
	current := start
	for loopCheck(current) {
		if dir == len(directions) {
			dir = 0
		}
		if directions[dir] == 'R' {
			current = paths[current][1]
		} else {
			current = paths[current][0]
		}
		dir++
		steps++
	}
	return steps
}

func lcm(numbers []int64) int64 {
	current := numbers[0]
	for i := 1; i < len(numbers); i++ {
		current = (current * numbers[i]) / gcd(current, numbers[i])
	}
	return current
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
