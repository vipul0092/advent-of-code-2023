package day16

import (
	"advent-of-code/utils"
	"fmt"
	"github.com/emirpasic/gods/queues/arrayqueue"
	"github.com/emirpasic/gods/sets/hashset"
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

type Pwd struct {
	p   Point
	dir Direction
}

var lines []string

func Solve() {
	lines = utils.ReadLines("/day16/input")

	fmt.Println("Part 1: ", getEnergisedCount(Pwd{Point{0, 0}, RIGHT}))

	maxi := 0
	for i := 0; i < len(lines); i++ {
		maxi = max(maxi, getEnergisedCount(Pwd{Point{i, 0}, RIGHT}))
		maxi = max(maxi, getEnergisedCount(Pwd{Point{i, len(lines[i]) - 1}, LEFT}))
	}
	for j := 0; j < len(lines[0]); j++ {
		maxi = max(maxi, getEnergisedCount(Pwd{Point{0, j}, DOWN}))
		maxi = max(maxi, getEnergisedCount(Pwd{Point{len(lines) - 1, j}, UP}))
	}

	fmt.Println("Part 2: ", maxi)
}

func getEnergisedCount(start Pwd) int {
	queue := arrayqueue.New()
	visited := hashset.New()
	energised := hashset.New()
	visited.Add(start)
	queue.Enqueue(start)

	for !queue.Empty() {
		curr, _ := queue.Dequeue()
		pos := curr.(Pwd).p
		direction := curr.(Pwd).dir
		energised.Add(pos)

		char := valueAt(pos)

		if char == '.' {
			enqueue(Pwd{move(pos, direction), direction}, queue, visited)
		} else if char == '\\' || char == '/' {
			var newdirection Direction
			if direction == UP || direction == DOWN {
				newdirection = rotate90(direction, char == '/')
			} else {
				newdirection = rotate90(direction, char == '\\')
			}
			enqueue(Pwd{move(pos, newdirection), newdirection}, queue, visited)
		} else if char == '|' {
			if direction == UP || direction == DOWN {
				enqueue(Pwd{move(pos, direction), direction}, queue, visited)
			} else {
				enqueue(Pwd{move(pos, UP), UP}, queue, visited)
				enqueue(Pwd{move(pos, DOWN), DOWN}, queue, visited)
			}
		} else if char == '-' {
			if direction == LEFT || direction == RIGHT {
				enqueue(Pwd{move(pos, direction), direction}, queue, visited)
			} else {
				enqueue(Pwd{move(pos, LEFT), LEFT}, queue, visited)
				enqueue(Pwd{move(pos, RIGHT), RIGHT}, queue, visited)
			}
		}
	}

	return energised.Size()
}

func enqueue(pwd Pwd, queue *arrayqueue.Queue, visited *hashset.Set) {
	if valid(pwd.p) && !visited.Contains(pwd) {
		queue.Enqueue(pwd)
		visited.Add(pwd)
	}
}

func rotate90(dir Direction, clockwise bool) Direction {
	switch dir {
	case UP:
		{
			if clockwise {
				return RIGHT
			}
			return LEFT
		}
	case DOWN:
		{
			if clockwise {
				return LEFT
			}
			return RIGHT
		}
	case LEFT:
		{
			if clockwise {
				return UP
			}
			return DOWN
		}
	case RIGHT:
		{
			if clockwise {
				return DOWN
			}
			return UP
		}
	default:
		panic("Impossible")
	}
}

func valid(p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(lines) && p.y < len(lines[0])
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
