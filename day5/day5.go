package day5

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Ranges struct {
	deststr string
	start   int64
	end     int64
	dest    int64
	count   int64
}

func Solve() {
	lines := utils.ReadLines("/day5/input")

	seeds := make([]int64, 0)
	rangeMap := make(map[string][]Ranges)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "seeds") {
			for _, s := range strings.Split(strings.Split(line, ": ")[1], " ") {
				seed, _ := strconv.ParseInt(s, 10, 64)
				seeds = append(seeds, seed)
			}
		} else if len(line) > 0 {
			srcdest := strings.Split(line[:len(line)-5], "-")
			i++
			ranges := make([]Ranges, 0)
			for i < len(lines) {
				line = lines[i]
				if len(line) == 0 {
					break
				}
				numstr := strings.Split(line, " ")
				start, _ := strconv.ParseInt(numstr[1], 10, 64)
				dest, _ := strconv.ParseInt(numstr[0], 10, 64)
				count, _ := strconv.ParseInt(numstr[2], 10, 64)

				ranges = append(ranges, Ranges{srcdest[2], start, start + count - 1, dest, count})
				i++
			}
			sort.Slice(ranges[:], func(i, j int) bool {
				if ranges[i].start == ranges[j].end {
					return ranges[i].end < ranges[j].end
				}
				return ranges[i].start < ranges[j].start
			})
			rangeMap[srcdest[0]] = ranges
		}
	}

	part1 := int64(math.MaxInt64)
	for _, seed := range seeds {
		part1 = min(part1, findMin(seed, seed, "seed", rangeMap))
	}
	fmt.Println("Part 1: ", part1)

	part2 := int64(math.MaxInt64)
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		count := seeds[i+1]

		part2 = min(part2, findMin(start, start+count-1, "seed", rangeMap))
	}
	fmt.Println("Part 2: ", part2)
}

func findMin(start int64, end int64, source string, rangeMap map[string][]Ranges) int64 {
	if source == "location" {
		return start
	}

	ranges := rangeMap[source]
	mininum := int64(math.MaxInt64)
	went := false
	deststr := ""

	for _, curr := range ranges {
		deststr = curr.deststr
		// * - [ - * - ]
		if start <= curr.start && end <= curr.end && end >= curr.start {
			mininum = min(mininum, findMin(mapp(curr.start, curr), mapp(end, curr), curr.deststr, rangeMap))
			went = true
			// [ - * - * - ]
		} else if start >= curr.start && end <= curr.end {
			mininum = min(mininum, findMin(mapp(start, curr), mapp(end, curr), curr.deststr, rangeMap))
			went = true
			// [ - * - ] - *
		} else if start >= curr.start && end >= curr.end && start <= curr.end {
			mininum = min(mininum, findMin(mapp(start, curr), mapp(curr.end, curr), curr.deststr, rangeMap))
			went = true
			// * - [ - ] - *
		} else if start <= curr.start && end >= curr.end {
			mininum = min(mininum, findMin(mapp(curr.start, curr), mapp(curr.end, curr), curr.deststr, rangeMap))
			went = true
		}
	}

	if !went {
		mininum = findMin(start, end, deststr, rangeMap)
	}
	return mininum
}

func mapp(val int64, rng Ranges) int64 {
	return rng.dest + (val - rng.start)
}
