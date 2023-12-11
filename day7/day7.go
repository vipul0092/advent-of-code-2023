package day7

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

var hand_priorities = make(map[[5]int]int)

func init() {
	hand_priorities[[5]int{1, 1, 1, 1, 1}] = 1
	hand_priorities[[5]int{2, 1, 1, 1, 0}] = 2
	hand_priorities[[5]int{2, 2, 1, 0, 0}] = 3
	hand_priorities[[5]int{3, 1, 1, 0, 0}] = 4
	hand_priorities[[5]int{3, 2, 0, 0, 0}] = 5
	hand_priorities[[5]int{4, 1, 0, 0, 0}] = 6
	hand_priorities[[5]int{5, 0, 0, 0, 0}] = 7
}

func Solve() {
	//lines := strings.Split(input, "\n")
	lines := utils.ReadLines("/day7/input")

	handmap := make(map[string]int)
	hands := make([]string, 0)

	for _, line := range lines {
		str := strings.Split(line, " ")
		num, _ := strconv.Atoi(str[1])
		handmap[str[0]] = num
		hands = append(hands, str[0])
	}

	fmt.Println("Part 1: ", sum(hands, handmap, false)) // 246409899
	fmt.Println("Part 2: ", sum(hands, handmap, true))  // 244848487
}

func sum(hands []string, handmap map[string]int, j bool) int {
	sort.Slice(hands[:], func(ii, jj int) bool {
		return compare(hands[ii], hands[jj], j)
	})

	total := 0
	for rank, h := range hands {
		total += (rank + 1) * handmap[h]
	}
	return total
}

func compare(s1, s2 string, j bool) bool {
	chars1 := make(map[rune]int, 5)
	chars2 := make(map[rune]int, 5)

	for i := 0; i < 5; i++ {
		chars1[rune(s1[i])]++
		chars2[rune(s2[i])]++
	}

	h1 := handPriority(chars1, j)
	h2 := handPriority(chars2, j)
	if h1 != h2 {
		return h1 < h2
	}

	for i := 0; i < 5; i++ {
		if s1[i] != s2[i] {
			c1 := cardPriority(rune(s1[i]), j)
			c2 := cardPriority(rune(s2[i]), j)
			return c1 < c2
		}
	}
	panic("not supposed to reach here")
}

func handPriority(counts map[rune]int, j bool) int {
	arr := [5]int{0, 0, 0, 0, 0}
	jcount := 0
	if j {
		jcount = counts['J']
		delete(counts, 'J')
	}
	idx := 0
	for _, count := range counts {
		arr[idx] = count
		idx++
	}
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] > arr[j]
	})
	if j {
		arr[0] += jcount
	}
	return hand_priorities[arr]
}

func cardPriority(card rune, j bool) int {
	priority := 0
	switch card {
	case 'A':
		priority = 13
	case 'K':
		priority = 12
	case 'Q':
		priority = 11
	case 'J':
		{
			if j {
				priority = 0
			} else {
				priority = 10
			}
		}
	case 'T':
		priority = 9
	default:
		priority = int(card - '0' - 1)
	}
	return priority
}
