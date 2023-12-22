package day22

import (
	"advent-of-code/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"sort"
	"strconv"
	"strings"
)

var space [501][11][11]int

type Point struct {
	x int
	y int
	z int
}

type Cube struct {
	p1 Point
	p2 Point
}

func Solve() {
	lines := utils.ReadLines("/day22/input")
	space = [501][11][11]int{}
	cubes := make([]Cube, len(lines))

	for i, line := range lines {
		p1str, p2str, _ := strings.Cut(line, "~")
		cubes[i] = Cube{getPoint(p1str), getPoint(p2str)}
	}

	sort.Slice(cubes, func(i, j int) bool {
		c1, c2 := cubes[i], cubes[j]
		if c1.p1.z == c1.p2.z && c2.p1.z == c2.p2.z {
			return c1.p1.z < c2.p2.z
		} else {
			return min(c1.p1.z, c1.p2.z) < min(c2.p1.z, c2.p2.z)
		}
	})

	for i, cube := range cubes {
		z1, z2 := getFirstBlockingZ(cube), -1
		if cube.p1.z != cube.p2.z { // vertical cube
			z2 = z1 + cube.p2.z - cube.p1.z
		} else {
			z2 = z1
		}
		updateSpace(i, z1, z2, cube)
	}

	supports, supportedBy := make(map[int]mapset.Set[int]), make(map[int]mapset.Set[int])
	for z := 1; z < 500; z++ {
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				// There is a brick here, check if something is above it and its different
				if space[z][x][y] != 0 && space[z+1][x][y] != 0 && space[z+1][x][y] != space[z][x][y] {
					// this means space[z][x][y] supports space[z+1][x][y]
					if _, has1 := supportedBy[space[z+1][x][y]]; !has1 {
						supportedBy[space[z+1][x][y]] = mapset.NewSet[int]()
					}
					supportedBy[space[z+1][x][y]].Add(space[z][x][y])

					if _, has2 := supports[space[z][x][y]]; !has2 {
						supports[space[z][x][y]] = mapset.NewSet[int]()
					}
					supports[space[z][x][y]].Add(space[z+1][x][y])
				}
			}
		}
	}

	singleSupporters := mapset.NewSet[int]()
	for _, supportList := range supportedBy {
		if supportList.Cardinality() == 1 {
			for val := range supportList.Iter() {
				singleSupporters.Add(val)
				break
			}
		}
	}
	fmt.Println("Part 1: ", len(cubes)-singleSupporters.Cardinality()) // 398

	total := 0
	for brick := range singleSupporters.Iter() {
		fallen := mapset.NewSet[int]()
		fallen.Add(brick)
		bricks := make([]int, 1)
		bricks[0] = brick

		for len(bricks) != 0 {
			newbricks := make([]int, 0)
			for _, current := range bricks {
				if supportedByCurrent, has := supports[current]; has {
					for supp := range supportedByCurrent.Iter() {
						haveAllFallen := supportedBy[supp].IsSubset(fallen) // all bricks that supports supp have fallen
						if haveAllFallen && !fallen.Contains(supp) {
							fallen.Add(supp)
							newbricks = append(newbricks, supp)
						}
					}
				}
			}
			bricks = newbricks
		}
		total += fallen.Cardinality() - 1
	}

	fmt.Println("Part 2: ", total) // 70727
}

func getFirstBlockingZ(cube Cube) int {
	z := cube.p1.z
	for z >= 1 {
		for x := cube.p1.x; x <= cube.p2.x; x++ {
			for y := cube.p1.y; y <= cube.p2.y; y++ {
				if space[z][x][y] != 0 {
					z++
					return z // Move one above because we found blocking at value `z`
				}
			}
		}
		z--
	}
	z++
	return z
}

func updateSpace(i, z1, z2 int, cube Cube) {
	x1, x2, y1, y2 := cube.p1.x, cube.p2.x, cube.p1.y, cube.p2.y
	for z := z1; z <= z2; z++ {
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				space[z][x][y] = i + 1
			}
		}
	}
}

func getPoint(pstr string) Point {
	pnums := strings.Split(pstr, ",")
	x, _ := strconv.Atoi(pnums[0])
	y, _ := strconv.Atoi(pnums[1])
	z, _ := strconv.Atoi(pnums[2])
	return Point{x, y, z}
}
