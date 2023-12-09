package day9

import (
	"advent-of-code/reader"
	"fmt"
	"strconv"
	"strings"
)

var part1Func = func(n int, ns []int) int {
	return n + ns[len(ns)-1]
}
var part2Func = func(n int, ns []int) int {
	return ns[0] - n
}

func Solve() {
	lines := reader.ReadLines("/day9/input")
	part1 := 0
	part2 := 0

	for _, line := range lines {
		nums := make([]int, 0)
		for _, ns := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(ns)
			nums = append(nums, num)
		}

		part1 += part1Func(getNum(nums, part1Func), nums)
		part2 += part2Func(getNum(nums, part2Func), nums)
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}

func getNum(nums []int, retcalc func(n int, ns []int) int) int {
	newnums := make([]int, len(nums)-1)
	allzero := true
	for i := 0; i < len(newnums); i++ {
		newnums[i] = nums[i+1] - nums[i]
		if newnums[i] != 0 {
			allzero = false
		}
	}

	if allzero {
		return 0
	}
	num := getNum(newnums, retcalc)
	return retcalc(num, newnums)
}
