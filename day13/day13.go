package day13

import (
	"advent-of-code/utils"
	"fmt"
)

func Solve() {
	lines := utils.ReadLines("/day13/input")

	sum, sum2 := 0, 0
	patterns := make([]string, 0)

	for i, line := range lines {
		if len(line) != 0 {
			patterns = append(patterns, line)
		}
		if len(line) == 0 || i == len(lines)-1 {
			sum += getScore(patterns, false)
			sum2 += getScore(patterns, true)
			patterns = make([]string, 0)
		}
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}

func getScore(patterns []string, smudge bool) int {
	vertical := false
	mid := getMid(getIntegers(patterns, false), smudge) // Horizontal
	if mid == -1 {
		vertical = true
		mid = getMid(getIntegers(patterns, true), smudge) // Vertical
	}
	if vertical {
		return mid + 1
	} else {
		return 100 * (mid + 1)
	}
}

func getMid(numbers []int, smudge bool) int {
	mid := -1
	for i := 0; i < len(numbers)-1; i++ {
		start := i
		smudgefound := false
		for end := i + 1; start >= 0 && end < len(numbers); end++ {
			xor := numbers[start] ^ numbers[end]
			// During exact match -> xor should be zero i.e. all bits match
			// During smudge match -> xor can be zero or a power of 2 i.e. one bit doesn't match
			valid := (xor == 0 && !smudge) || (smudge && (xor&(xor-1)) == 0)
			smudgefound = smudgefound || (smudge && valid && xor > 0)
			if valid {
				mid = i
			} else {
				mid = -1
				break
			}
			start--
		}
		if mid != -1 && (!smudge || smudgefound) {
			break
		}
		mid = -1
	}
	return mid
}

func getIntegers(patterns []string, vertical bool) []int {
	ints := make([]int, 0)
	if !vertical {
		for _, pattern := range patterns {
			num := 0
			for i, char := range pattern {
				if char == '#' {
					num += 1 << i
				}
			}
			ints = append(ints, num)
		}
	} else {
		for j := 0; j < len(patterns[0]); j++ {
			num := 0
			for i, pattern := range patterns {
				if pattern[j] == '#' { // consider # as 1
					num += 1 << i
				}
			}
			ints = append(ints, num)
		}
	}
	return ints
}
