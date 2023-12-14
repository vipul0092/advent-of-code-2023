package day14

import (
	"advent-of-code/utils"
	"fmt"
	"github.com/emirpasic/gods/stacks"
	"github.com/emirpasic/gods/stacks/arraystack"
	"strings"
)

var dimension int

type Direction int

const MaxCycles = 1000000000
const (
	NORTH Direction = iota + 1
	WEST
	SOUTH
	EAST
)

type Rock struct {
	ch  rune
	loc int
}

func Solve() {
	lines := utils.ReadLines("/day14/input")
	dimension = len(lines)
	grid := createGrid()

	for i, line := range lines {
		for j, char := range line {
			grid[i][j] = char
		}
	}

	part1 := getNorthLoad(tilt(grid, NORTH))
	fmt.Println("Part 1: ", part1) // 107053

	// ---------- Part 2 Begins ---------
	mapping := make(map[string]int)
	mapping[getHash(grid)] = 0

	subtract, cycle := 0, -1
	for i := 1; i <= MaxCycles; i++ {
		newgrid := doCycle(grid)
		hash := getHash(newgrid)
		grid = newgrid

		previndex, has := mapping[hash]
		if has {
			subtract = previndex
			cycle = i - previndex
			fmt.Println("Found cycle at iteration: ", i)
			break
		} else {
			mapping[hash] = i
		}
	}

	todo := (MaxCycles - subtract) % cycle
	for todo > 0 {
		todo--
		grid = doCycle(grid)
	}
	part2 := getNorthLoad(grid)
	fmt.Println("Part 2: ", part2) // 88371
}

func doCycle(grid [][]rune) [][]rune {
	grid = tilt(grid, NORTH)
	grid = tilt(grid, WEST)
	grid = tilt(grid, SOUTH)
	grid = tilt(grid, EAST)
	return grid
}

func tilt(grid [][]rune, direction Direction) [][]rune {
	allstacks := make(map[int]stacks.Stack)
	vertical, north, west := direction == NORTH || direction == SOUTH, direction == NORTH, direction == WEST

	for i := 0; i < dimension; i++ {
		allstacks[i] = arraystack.New()
		for j, k := 0, dimension-1; j < dimension; j++ {
			idx := -1
			if north || west {
				idx = j
			} else {
				idx = k
			}
			ch := '0'
			if vertical {
				ch = grid[idx][i]
			} else {
				ch = grid[i][idx]
			}

			if ch == '#' {
				allstacks[i].Push(Rock{ch, idx})
			} else if ch == 'O' {
				pidx := -1
				if allstacks[i].Size() == 0 {
					if north || west {
						pidx = 0
					} else {
						pidx = dimension - 1
					}
				} else {
					top, _ := allstacks[i].Peek()
					pidx = top.(Rock).loc
					if north || west {
						pidx++
					} else {
						pidx--
					}
				}
				allstacks[i].Push(Rock{ch, pidx})
			}
			k--
		}
	}

	newgrid := createGrid()
	for i := 0; i < dimension; i++ {
		stack := allstacks[i]
		for stack.Size() > 0 {
			rock, _ := stack.Pop()
			if vertical {
				newgrid[rock.(Rock).loc][i] = rock.(Rock).ch
			} else {
				newgrid[i][rock.(Rock).loc] = rock.(Rock).ch
			}
		}
		for j := 0; j < dimension; j++ {
			ch := '0'
			if vertical {
				ch = newgrid[i][j]
			} else {
				ch = newgrid[j][i]
			}

			if ch != 'O' && ch != '#' {
				if vertical {
					newgrid[i][j] = '.'
				} else {
					newgrid[j][i] = '.'
				}
			}
		}
	}
	return newgrid
}

func getHash(grid [][]rune) string {
	var sb strings.Builder
	for _, row := range grid {
		for _, char := range row {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}

func getNorthLoad(grid [][]rune) int {
	load := 0
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if grid[i][j] == 'O' {
				load += dimension - i
			}
		}
	}
	return load
}

func createGrid() [][]rune {
	grid := make([][]rune, dimension)
	for i := 0; i < dimension; i++ {
		grid[i] = make([]rune, dimension)
	}
	return grid
}
