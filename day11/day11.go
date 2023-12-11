package day11

import (
	"advent-of-code/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
)

type Point struct {
	x int
	y int
}

func Solve() {
	lines := utils.ReadLines("/day11/input")
	norows := mapset.NewSet[int]()
	nocols := mapset.NewSet[int]()
	galaxies := make([]Point, 0)

	for i, line := range lines {
		blank := true
		for j, char := range line {
			if char == '#' {
				blank = false
				galaxies = append(galaxies, Point{i, j})
			}
		}
		if blank {
			norows.Add(i)
		}
	}

	for j := 0; j < len(lines[0]); j++ {
		blank := true
		for i := 0; i < len(lines); i++ {
			if lines[i][j] == '#' {
				blank = false
				break
			}
		}
		if blank {
			nocols.Add(j)
		}
	}

	fmt.Println("Part 1: ", getPathSum(galaxies, norows, nocols, 2))
	fmt.Println("Part 2: ", getPathSum(galaxies, norows, nocols, 1000000))
}

func getPathSum(galaxies []Point, norows, nocols mapset.Set[int], times int) int {
	total := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			g1 := galaxies[i]
			g2 := galaxies[j]

			minx := min(g1.x, g2.x)
			maxx := max(g1.x, g2.x)
			miny := min(g1.y, g2.y)
			maxy := max(g1.y, g2.y)
			diff := maxx - minx + maxy - miny

			for k := minx + 1; k < maxx; k++ {
				if norows.Contains(k) {
					diff += times - 1
				}
			}
			for k := miny + 1; k < maxy; k++ {
				if nocols.Contains(k) {
					diff += times - 1
				}
			}

			total += diff
		}
	}
	return total
}
