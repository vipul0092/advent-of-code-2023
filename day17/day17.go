package day17

import (
	"advent-of-code/utils"
	"fmt"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"maps"
	"math"
)

var adjacencyList map[Pwd]map[Pwd]int
var distanceMap map[Pwd]int

type Direction int

const (
	HORIZONTAL Direction = iota
	VERTICAL
)

type Pair struct {
	left  int
	right Pwd
}
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
	lines = utils.ReadLines("/day17/input")
	start := Point{0, 0}
	target := Point{len(lines) - 1, len(lines[0]) - 1}

	populateGraph(0, 3)
	distanceCopy := make(map[Pwd]int)
	maps.Copy(distanceCopy, distanceMap)
	min1 := dijkstra(Pwd{start, HORIZONTAL}, target)
	distanceMap = make(map[Pwd]int)
	maps.Copy(distanceMap, distanceCopy)
	min2 := dijkstra(Pwd{start, VERTICAL}, target)
	fmt.Println("Part 1: ", min(min1, min2))

	populateGraph(3, 10)
	distanceCopy = make(map[Pwd]int)
	maps.Copy(distanceCopy, distanceMap)
	min1 = dijkstra(Pwd{start, HORIZONTAL}, target)
	distanceMap = make(map[Pwd]int)
	maps.Copy(distanceMap, distanceCopy)
	min2 = dijkstra(Pwd{start, VERTICAL}, target)
	fmt.Println("Part 2: ", min(min1, min2))
}

func populateGraph(mindis, maxdis int) {
	adjacencyList = make(map[Pwd]map[Pwd]int)
	distanceMap = make(map[Pwd]int)

	for i, line := range lines {
		for j := range line {
			populateNeighbors(i, j, mindis, maxdis, HORIZONTAL, VERTICAL, func(i, j, k int, subtract bool) (int, int) {
				if subtract {
					return i + k, j
				} else {
					return i - k, j
				}
			})
			populateNeighbors(i, j, mindis, maxdis, VERTICAL, HORIZONTAL, func(i, j, k int, subtract bool) (int, int) {
				if subtract {
					return i, j + k
				} else {
					return i, j - k
				}
			})
		}
	}
}

func populateNeighbors(i, j, mindis, maxdis int, pointDir Direction, neighborDir Direction,
	mutator func(i, j, k int, subtract bool) (int, int)) {
	p := Pwd{Point{i, j}, pointDir}
	adjacencyList[p] = make(map[Pwd]int)
	distanceMap[p] = math.MaxInt64

	distancep, distancen := 0, 0
	for k := 1; k <= maxdis; k++ {
		ii, jj := mutator(i, j, k, false)
		if valid(ii, jj) {
			distancep += valueAt(ii, jj)
			if k > mindis {
				adjacencyList[p][Pwd{Point{ii, jj}, neighborDir}] = distancep
			}
		}

		ii, jj = mutator(i, j, k, true)
		if valid(ii, jj) {
			distancen += valueAt(ii, jj)
			if k > mindis {
				adjacencyList[p][Pwd{Point{ii, jj}, neighborDir}] = distancen
			}
		}
	}
}

func dijkstra(start Pwd, end Point) int {
	length, current := -1, start
	distanceMap[start] = 0

	minheap := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(Pair).left - b.(Pair).left
	})
	minheap.Enqueue(Pair{0, start})
	visited := [150][150][2]bool{}

	for !minheap.Empty() {
		top, _ := minheap.Dequeue()
		current = top.(Pair).right
		visited[current.p.x][current.p.y][current.dir] = true
		currentDistance := top.(Pair).left

		if current.p == end {
			length = currentDistance
			break
		}

		for neighbor, dis := range adjacencyList[current] {
			p := neighbor.p

			if visited[p.x][p.y][neighbor.dir] {
				continue
			}
			neighbourNewDistance := currentDistance + dis
			neighborCurrentDistance := distanceMap[neighbor]

			// assign the distance to the neighbour
			// and update the distance in unvisited map
			if neighbourNewDistance < neighborCurrentDistance {
				distanceMap[neighbor] = neighbourNewDistance
				minheap.Enqueue(Pair{neighbourNewDistance, neighbor})
			}
		}
	}
	return length
}

func valueAt(i, j int) int {
	return int(lines[i][j] - '0')
}

func valid(i, j int) bool {
	return i >= 0 && j >= 0 && i < len(lines) && j < len(lines[0])
}
