package day10

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

type Point struct {
	x int
	y int
}

var DIRECTIONS []Direction
var INITIAL_MOVE map[Direction]mapset.Set[rune]
var CAN_MOVE map[rune]mapset.Set[Direction]
var lines []string

func init() {
	DIRECTIONS = []Direction{UP, DOWN, LEFT, RIGHT}
	INITIAL_MOVE = make(map[Direction]mapset.Set[rune])
	INITIAL_MOVE[UP] = mapset.NewSet('|', '7', 'F')
	INITIAL_MOVE[DOWN] = mapset.NewSet('|', 'L', 'J')
	INITIAL_MOVE[LEFT] = mapset.NewSet('-', 'F', 'L')
	INITIAL_MOVE[RIGHT] = mapset.NewSet('-', '7', 'J')

	CAN_MOVE = make(map[rune]mapset.Set[Direction])
	CAN_MOVE['|'] = mapset.NewSet(UP, DOWN)
	CAN_MOVE['-'] = mapset.NewSet(LEFT, RIGHT)
	CAN_MOVE['F'] = mapset.NewSet(DOWN, RIGHT)
	CAN_MOVE['7'] = mapset.NewSet(LEFT, DOWN)
	CAN_MOVE['L'] = mapset.NewSet(UP, RIGHT)
	CAN_MOVE['J'] = mapset.NewSet(UP, LEFT)
}

func Solve() {
	lines = utils.ReadLines("/day10/input")
	var spos Point
	allpoints := make([]Point, 0)
	for i, line := range lines {
		for j, ch := range line {
			pt := Point{i, j}
			allpoints = append(allpoints, pt)
			if ch == 'S' {
				spos = pt
			}
		}
	}

	startPoint, direction := getStartingPoint(spos)

	prev := spos
	vertices := make([]Point, 0)
	polygonPoints := mapset.NewSet[Point]()
	current := startPoint
	vertices = append(vertices, spos)
	polygonPoints.Add(current)
	totalDistance := 0

	for valueAt(current) != 'S' {
		curr := valueAt(current)
		for nextDirection := range CAN_MOVE[curr].Iter() {
			moved := move(current, nextDirection)
			if valid(moved) && moved != prev {
				polygonPoints.Add(moved)
				prev = current
				current = moved
				if direction != nextDirection { // Add a vertex when changing direction
					direction = nextDirection
					vertices = append(vertices, moved)
				}
				totalDistance++
				break
			}
		}
	}

	maxi := 0
	if totalDistance%2 == 0 {
		maxi = totalDistance / 2
	} else {
		maxi = (totalDistance / 2) + 1
	}

	inside := 0
	for _, pt := range allpoints {
		if !polygonPoints.Contains(pt) && pointInPolygon(pt, vertices) {
			inside++
		}
	}

	fmt.Println("Part 1: ", maxi)
	fmt.Println("Part 2: ", inside)
}

// Implementation Converted to Go
// Java Source: https://www.codingninjas.com/studio/library/check-if-a-point-lies-in-the-interior-of-a-polygon
func pointInPolygon(point Point, polygon []Point) bool {
	numVertices := len(polygon)
	x := point.x
	y := point.y
	inside := false
	p1 := polygon[0]

	for i := 1; i <= numVertices; i++ {
		p2 := polygon[i%numVertices]
		if y > min(p1.y, p2.y) {
			if y <= max(p1.y, p2.y) {
				if x <= max(p1.x, p2.x) {
					xIntersection := float32(y-p1.y)*float32(p2.x-p1.x)/float32(p2.y-p1.y) + float32(p1.x)
					if p1.x == p2.x || float32(x) <= xIntersection {
						inside = !inside
					}
				}
			}
		}
		p1 = p2
	}
	return inside
}

func getStartingPoint(spos Point) (Point, Direction) {
	for _, dir := range DIRECTIONS {
		moved := move(spos, dir)
		if valid(moved) && INITIAL_MOVE[dir].Contains(valueAt(moved)) {
			return moved, dir
		}
	}
	panic("wont reach here")
}

func valid(p Point) bool {
	return !(p.x < 0 || p.y < 0 || p.x >= len(lines) || p.y >= len(lines[0])) && valueAt(p) != '.'
}

func valueAt(p Point) rune {
	return rune(lines[p.x][p.y])
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
