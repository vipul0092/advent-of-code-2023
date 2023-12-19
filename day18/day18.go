package day18

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Direction rune

const (
	UP    Direction = 'U'
	DOWN  Direction = 'D'
	LEFT  Direction = 'L'
	RIGHT Direction = 'R'
)

type Point struct {
	x int
	y int
}

func Solve() {
	lines := utils.ReadLines("/day18/input")

	current, vertices, total := Point{0, 0}, make([]Point, 0), 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir := Direction(parts[0][0])
		dis, _ := strconv.Atoi(parts[1])
		current = move(current, dis, dir)
		total += dis
		vertices = append(vertices, current)
	}
	area := getArea(vertices)
	area += total - ((total / 2) - 1)
	fmt.Println("Part 1: ", area)

	current, vertices, total = Point{0, 0}, make([]Point, 0), 0
	for _, line := range lines {
		hex := strings.Split(line, " ")[2]
		hex = hex[2 : len(hex)-1]
		dir := getDirection(hex)
		dis, _ := strconv.ParseInt(hex[:len(hex)-1], 16, 64)
		current = move(current, int(dis), dir)
		total += int(dis)
		vertices = append(vertices, current)
	}
	area = getArea(vertices)
	area += total - ((total / 2) - 1)
	fmt.Println("Part 2: ", area)
}

func getArea(poly []Point) int {
	res := 0.0
	for i := 0; i < len(poly); i++ {
		var p Point
		if i > 0 {
			p = poly[i-1]
		} else {
			p = poly[len(poly)-1]
		}
		var q = poly[i]
		res += float64(p.x-q.x) * float64(p.y+q.y)
	}
	res = math.Abs(res) / 2
	return int(res)
}

func move(p Point, dis int, dir Direction) Point {
	switch dir {
	case UP:
		return Point{p.x, p.y + dis}
	case DOWN:
		return Point{p.x, p.y - dis}
	case LEFT:
		return Point{p.x - dis, p.y}
	case RIGHT:
		return Point{p.x + dis, p.y}
	default:
		panic("Impossible")
	}
}

func getDirection(hex string) Direction {
	diri := int(hex[len(hex)-1] - '0')
	switch diri {
	case 0:
		return RIGHT
	case 1:
		return DOWN
	case 2:
		return LEFT
	case 3:
		return UP
	default:
		panic("Impossible")
	}
}
