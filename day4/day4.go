package day4

import (
	"advent-of-code/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"strconv"
	"strings"
)

const input = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func Solve() {
	//lines := strings.Split(input, "\n")
	lines := strings.Split(utils.Read("/day4/input"), "\n")

	sum := 0
	counts := make(map[int]int)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		cardstr := strings.Split(parts[0], " ")
		card, _ := strconv.Atoi(cardstr[len(cardstr)-1])

		counts[card]++

		numstr := strings.Split(parts[1], " | ")
		winstr := strings.Split(numstr[0], " ")
		havestr := strings.Split(numstr[1], " ")

		wins := mapset.NewSet[int]()
		for _, w := range winstr {
			if len(w) != 0 {
				tmp, _ := strconv.Atoi(w)
				wins.Add(tmp)
			}
		}

		score := 0
		found := 0
		for _, h := range havestr {
			if len(h) != 0 {
				tmp, _ := strconv.Atoi(h)
				if wins.Contains(tmp) {
					found++
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
				}
			}
		}

		for i := 1; i <= found; i++ {
			counts[card+i] += counts[card]
		}
		sum += score
	}

	fmt.Println("Part 1: ", sum)

	sum2 := 0
	for _, cnt := range counts {
		sum2 += cnt
	}
	fmt.Println("Part 2: ", sum2)
}
