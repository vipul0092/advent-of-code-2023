package day3

import (
	"advent-of-code/reader"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"strings"
)

const input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var diffs = []int{-1, 0, 1}

type Point struct {
	x int
	y int
}

type Num struct {
	number     int
	neighbours mapset.Set[Point]
}

func Solve() {
	//lines := strings.Split(input, "\n")
	lines := strings.Split(reader.Read("/day3/input"), "\n")

	numsmap := make(map[Point]int)
	numbers := make([]Num, 0)
	gears := make([]Point, 0)
	symbols := mapset.NewSet[Point]()

	for lidx, line := range lines {
		num := 0
		neighbours := mapset.NewSet[Point]()
		self := make([]Point, 0)
		for cidx, char := range line {
			if char >= '0' && char <= '9' {
				num = (num * 10) + int(char-'0')
				for _, dx := range diffs {
					for _, dy := range diffs {
						if dx == 0 && dy == 0 {
							continue
						}
						neighbours.Add(Point{lidx + dx, cidx + dy})
					}
				}
				self = append(self, Point{lidx, cidx})
			} else {
				if char != '.' {
					symbols.Add(Point{lidx, cidx})
				}
				if char == '*' {
					gears = append(gears, Point{lidx, cidx})
				}
				if num != 0 {
					numbers = append(numbers, Num{num, neighbours})
					for _, s := range self {
						numsmap[s] = num
					}
				}

				num = 0
				neighbours = mapset.NewSet[Point]()
				self = make([]Point, 0)
			}
		}

		if num != 0 {
			numbers = append(numbers, Num{num, neighbours})
			for _, s := range self {
				numsmap[s] = num
			}
		}
	}

	sum := 0
	for _, number := range numbers {
		valid := false
		for nbor := range number.neighbours.Iter() {
			if symbols.Contains(nbor) {
				valid = true
				break
			}
		}

		if valid {
			sum += number.number
		}
	}

	fmt.Println("Part 1: ", sum)

	// Part 2 begins
	sum2 := 0
	for _, gear := range gears {
		found := mapset.NewSet[int]()
		for _, dx := range diffs {
			for _, dy := range diffs {
				if dx == 0 && dy == 0 {
					continue
				}
				tmp := Point{gear.x + dx, gear.y + dy}
				fnd, has := numsmap[tmp]
				if has {
					found.Add(fnd)
				}
			}
		}

		if found.Cardinality() == 2 {
			product := 1
			for f := range found.Iter() {
				product *= f
			}
			sum2 += product
		}
	}

	fmt.Println("Part 2: ", sum2)
}
