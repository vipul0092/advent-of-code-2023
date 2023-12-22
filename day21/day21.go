package day21

import (
	"advent-of-code/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var DIRECTIONS = [4]Direction{UP, DOWN, LEFT, RIGHT}

type Point struct {
	x int
	y int
}

type Pwd struct {
	p   Point
	dir Direction
}

var lines []string

func Solve() {
	lines = utils.ReadLines("/day21/input")
	neighbors := make(map[Point][]Point)
	loopedNeighbors := make(map[Point][]Pwd)
	var start Point

	for i, line := range lines {
		for j, char := range line {
			if char == 'S' || char == '.' {
				p := Point{i, j}
				if char == 'S' {
					start = p
				}
				neighbors[p] = make([]Point, 0)
				loopedNeighbors[p] = make([]Pwd, 0)

				for _, dir := range DIRECTIONS {
					pn := move(p, dir)
					if !valid(pn) { // lies outside the grid, so loop it back
						looped := getLoopedPoint(pn)
						// no need to check validSign because looped points are always valid
						loopedNeighbors[p] = append(loopedNeighbors[p], looped)
					} else if valid(pn) && validSign(pn) {
						neighbors[p] = append(neighbors[p], pn)
					}
				}
			}
		}
	}

	fmt.Println("Part 1: ",
		getStepPoints(64, start, mapset.NewSet[int](64), neighbors, loopedNeighbors)[64]) // 3605

	// Explanation for Quadratic solution
	// The no. of steps to calculate in part 2 is 26501365 = (202300*131) + 65
	// If the no. of points after 131*X + 65 steps is f(X), then f(X) is quadratic
	// This is hinted via the example as well
	// see https://old.reddit.com/r/adventofcode/comments/18nevo3/2023_day_21_solutions/keaiiq7/ as to why
	//
	// To calc final answer, we can put X = 202300
	// now f(X) = aX^2 + bX + c, f(X) gives number of pts after 65, 65 + 131, 65 + 131*2 steps and so on
	// To find the quadratic, we have evaluate the f(X) at X=0,1,2
	// f(0) = pts after 65 steps, f(1) = ptr after 65 + 131 steps, f(2) = 65 + 131*2 steps
	// Using these 3 points, we can calculate a, b & c in the quadratic
	// and hence calculate f(202300) which is the answer
	zero, one, two := 65, 65+131, 65+(131*2)
	points := getStepPoints(65+(131*2), start, mapset.NewSet[int](zero, one, two), neighbors, loopedNeighbors)
	f0, f1, f2 := points[zero], points[one], points[two]

	// f(X) = aX^2 + bX + c
	// f(0) = c; f(1) = a + b + c; f(2) = 4a + 2b + c
	// a = ((f(2) - c)/2) - (f(1) - c)
	// b = f(1) - c - a
	c := f0
	a := ((f2 - c) / 2) - (f1 - c)
	b := f1 - c - a

	X := 202300
	answer := (a * X * X) + (b * X) + c // f(202300)
	fmt.Println("Part 2: ", answer)     // 596734624269210
}

func getStepPoints(totalSteps int, start Point, stepsToCapture mapset.Set[int],
	neighbors map[Point][]Point, loopedNeighbors map[Point][]Pwd) map[int]int {
	retval := make(map[int]int)
	pointsPerGrid := make(map[Point]mapset.Set[Point])
	grid := Point{0, 0}

	// Always start from the base
	pointsPerGrid[grid] = mapset.NewSet[Point]()
	pointsPerGrid[grid].Add(start)

	steps := 0
	for {
		newGridPoints := make(map[Point]mapset.Set[Point])
		for gr, pt := range pointsPerGrid {
			gridPos, gridPoints := gr, pt
			if _, has := newGridPoints[gridPos]; !has {
				newGridPoints[gridPos] = mapset.NewSet[Point]()
			}
			for p := range gridPoints.Iter() {
				for _, n := range neighbors[p] {
					newGridPoints[gridPos].Add(n)
				}

				for _, n := range loopedNeighbors[p] {
					newgrid := move(gridPos, n.dir)
					if _, has := newGridPoints[newgrid]; !has {
						newGridPoints[newgrid] = mapset.NewSet[Point]()
					}
					newGridPoints[newgrid].Add(n.p)
				}
			}

		}
		steps++
		pointsPerGrid = newGridPoints
		if stepsToCapture.Contains(steps) {
			retval[steps] = getPointsCount(pointsPerGrid)
		}
		if steps == totalSteps {
			break
		}
	}
	return retval
}

func getLoopedPoint(p Point) Pwd {
	if p.x < 0 {
		return Pwd{Point{len(lines) - 1, p.y}, UP} // loop back down
	}
	if p.y < 0 {
		return Pwd{Point{p.x, len(lines[0]) - 1}, LEFT} // loop back right
	}
	if p.x >= len(lines) {
		return Pwd{Point{0, p.y}, DOWN} // loop back up
	}
	if p.y >= len(lines[0]) {
		return Pwd{Point{p.x, 0}, RIGHT} // loop back left
	}
	panic("Impossible")
}

func getPointsCount(pointsPerGrid map[Point]mapset.Set[Point]) int {
	total := 0
	for _, pts := range pointsPerGrid {
		total += pts.Cardinality()
	}
	return total
}

func valid(p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(lines) && p.y < len(lines[0])
}

func validSign(p Point) bool {
	return valueAt(p.x, p.y) == '.' || valueAt(p.x, p.y) == 'S'
}

func valueAt(x, y int) rune {
	return rune(lines[x][y])
}

func move(p Point, direction Direction) Point {
	switch direction {
	case UP:
		return Point{p.x - 1, p.y}
	case DOWN:
		return Point{p.x + 1, p.y}
	case LEFT:
		return Point{p.x, p.y - 1}
	case RIGHT:
		return Point{p.x, p.y + 1}
	default:
		panic("Not possible")
	}
}
