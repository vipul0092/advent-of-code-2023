package day23

import (
	"advent-of-code/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/emirpasic/gods/queues/arrayqueue"
	"time"
)

var lines []string

type Point struct {
	i int
	j int
}

type Pwl struct {
	p   Point
	len int
}

type DP [142][142][142][142]int32
type GRAPH map[int]map[int]int

var graph GRAPH

var dp DP

var diffs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func SolvePart1() {
	lines = utils.ReadLines("/day23/input")
	start := time.Now()
	fmt.Println("Part 1: ", dfs(0, 1, 0, 0, len(lines)-1, len(lines[0])-2)) // 2134
	fmt.Println("Took: ", time.Since(start))
}

func SolvePart2() {
	lines = utils.ReadLines("/day23/input")
	start := time.Now()
	rows, cols, id := len(lines), len(lines[0]), 3

	adjacencyList := make(map[Point][]Point)

	compressedGraphNodes := make(map[Point]int)
	compressedGraphNodes[Point{0, 1}] = 1
	compressedGraphNodes[Point{rows - 1, cols - 2}] = 2
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !valid(i, j) {
				continue
			}
			pt := Point{i, j}
			adjacencyList[pt] = make([]Point, 0)
			for _, df := range diffs {
				ni, nj := i+df[0], j+df[1]
				if valid(ni, nj) {
					adjacencyList[pt] = append(adjacencyList[pt], Point{ni, nj})
				}
			}

			if len(adjacencyList[pt]) > 2 {
				compressedGraphNodes[pt] = id
				id++
			}
		}
	}

	graph = GRAPH{}
	for pt := range compressedGraphNodes {
		populateGraphForNode(pt, compressedGraphNodes, adjacencyList)
	}
	fmt.Println("Part 2: ", dfs2(1, 1, 2, 0)) // 6298
	fmt.Println("Took: ", time.Since(start))
}

func dfs2(current, prev, end int, visited int) int {
	if current == end {
		return 0
	}
	maxi := -1
	for neighbor, distance := range graph[current] {
		if neighbor != prev && !isBitSet(visited, neighbor) {
			d := dfs2(neighbor, current, end, setBit(visited, neighbor))
			if d != -1 {
				maxi = max(maxi, d+distance)
			}
		}
	}
	return maxi
}

func dfs(ci, cj, pi, pj, ei, ej int) int {
	if ci == ei && cj == ej {
		return 0
	}

	if dp[ci][cj][pi][pj] != 0 {
		return int(dp[ci][cj][pi][pj] - 2)
	}

	maxi := -1
	for _, df := range diffs {
		ni, nj := ci+df[0], cj+df[1]
		if !(ni == pi && nj == pj) && valid(ni, nj) {
			if lines[ni][nj] != '.' {
				if lines[ni][nj] == '>' {
					nj++
				} else if lines[ni][nj] == '<' {
					nj--
				} else if lines[ni][nj] == '^' {
					ni--
				} else {
					ni++
				}
			}
			if ni == pi && nj == pj {
				continue
			}
			if d := dfs(ni, nj, ci, cj, ei, ej); d != -1 {
				maxi = max(maxi, d+1)
			}
		}
	}
	dp[ci][cj][pi][pj] = int32(maxi + 2)
	return maxi
}

func populateGraphForNode(point Point, compressedGraphNodes map[Point]int, adjacencyList map[Point][]Point) {
	id, queue, visited := compressedGraphNodes[point], arrayqueue.New(), mapset.NewSet[Point]()
	queue.Enqueue(Pwl{point, 1})
	visited.Add(point)

	for !queue.Empty() {
		pwl, _ := queue.Dequeue()
		p, distance := pwl.(Pwl).p, pwl.(Pwl).len
		for _, neighbor := range adjacencyList[p] {
			if !visited.Contains(neighbor) {
				if nid, has := compressedGraphNodes[neighbor]; has {
					if _, hid := graph[id]; !hid {
						graph[id] = make(map[int]int)
					}
					if _, hnid := graph[nid]; !hnid {
						graph[nid] = make(map[int]int)
					}
					graph[id][nid] = distance
					graph[nid][id] = distance
				} else {
					queue.Enqueue(Pwl{neighbor, distance + 1})
				}
				visited.Add(neighbor)
			}
		}
	}
}

func valid(i, j int) bool {
	return i >= 0 && j >= 0 && i < len(lines) && j < len(lines[0]) && lines[i][j] != '#'
}

func isBitSet(bitSet, bitPos int) bool {
	return (bitSet & (1 << bitPos)) != 0
}

func setBit(bitSet, bitPos int) int {
	return bitSet | (1 << bitPos)
}
