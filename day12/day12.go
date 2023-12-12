package day12

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

var dp [104][104][104]int

const input = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func Solve() {
	//lines := strings.Split(input, "\n")
	lines := utils.ReadLines("/day12/input")

	sum := 0
	sum2 := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		numbers := make([]int, 0)
		for _, nstr := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(nstr)
			numbers = append(numbers, num)
		}
		sum += getAnswer(1, numbers, parts[0])
		sum2 += getAnswer(5, numbers, parts[0])
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}

func getAnswer(times int, numbers []int, pattern string) int {
	allnums := make([]int, 0)
	var sb strings.Builder
	for times > 0 {
		times--
		sb.WriteString(pattern)
		if times != 0 {
			sb.WriteRune('?')
		}
		allnums = append(allnums, numbers...)
	}
	dp = [104][104][104]int{}
	fullpattern := sb.String()
	return recurse(fullpattern, 0, allnums, 0, 0)
}

func recurse(pattern string, pidx int, numbers []int, nidx int, grouplen int) int {
	if len(pattern) == pidx {
		if (nidx == len(numbers)-1 && numbers[nidx] == grouplen) || (nidx == len(numbers) && grouplen == 0) {
			return 1
		}
		return 0
	}

	if dp[pidx][nidx][grouplen] != 0 {
		return dp[pidx][nidx][grouplen] - 1
	}
	sum := 0
	char := pattern[pidx]

	if char == '?' || char == '#' {
		// place a '#' and increment the grouplen
		sum += recurse(pattern, pidx+1, numbers, nidx, grouplen+1)
	}
	if char == '?' || char == '.' {
		// if grouplen > 0, we can place a '.' and close the group if the count matches
		if grouplen > 0 && nidx < len(numbers) && numbers[nidx] == grouplen {
			sum += recurse(pattern, pidx+1, numbers, nidx+1, 0)
		}
		// if no group, place a '.' and simply move ahead without any matching
		if grouplen == 0 {
			sum += recurse(pattern, pidx+1, numbers, nidx, 0)
		}
	}

	dp[pidx][nidx][grouplen] = sum + 1
	return sum
}
