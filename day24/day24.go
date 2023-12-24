package day24

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

const LEAST = 200000000000000
const MOST = 400000000000000

// Line represents a 3-D line as ax + by + c = 0
type Line struct {
	a float64
	b float64
	c float64
}

type Stone struct {
	x    int
	y    int
	z    int
	vx   int
	vy   int
	vz   int
	line Line
}

func Solve() {
	lines := utils.ReadLines("/day24/input")

	stones := make([]Stone, len(lines))
	for i, line := range lines {
		p, v, _ := strings.Cut(line, " @ ")
		x, y, z := parseNums(p)
		vx, vy, vz := parseNums(v)

		m := float64(vy) / float64(vx)     // y2-y1 / x2-x1
		c := float64(y) - (m * float64(x)) // y = mx + c => c = y - mx
		// y = mx + c => mx - y + c = 0
		stones[i] = Stone{x, y, z, vx, vy, vz, Line{m, -1, c}}
	}

	count := 0
	for i := 0; i < len(stones)-1; i++ {
		for j := i + 1; j < len(stones); j++ {
			s1, s2 := stones[i], stones[j]
			// intersection of two lines a1x + b1y + c1 = 0 & a2x + b2y + c2 = 0 is:
			// b1c2 - b2c1 / a1b2 - a2b1, c1a2 - c2a1 / a1b2 - a2b1
			a1, b1, c1, a2, b2, c2 := s1.line.a, s1.line.b, s1.line.c, s2.line.a, s2.line.b, s2.line.c
			a1b2_a2b1 := (a1 * b2) - (a2 * b1)
			if a1b2_a2b1 == 0 { // parallel or same
				continue
			}

			ix, iy := (b1*c2-b2*c1)/a1b2_a2b1, (c1*a2-c2*a1)/a1b2_a2b1

			// time should be > 0
			// x(t) = x0 + vt => t = x(t) - x0 / v
			t1, t2 := (ix-float64(s1.x))/float64(s1.vx), (ix-float64(s2.x))/float64(s2.vx)

			if t1 > 0 && t2 > 0 && ix >= LEAST && ix <= MOST && iy >= LEAST && iy <= MOST {
				count++
			}
		}
	}

	fmt.Println("Part 1: ", count) // 13892

	// Generate SageMath script
	fmt.Println()
	fmt.Println("var('x y z vx vy vz t1 t2 t3')")
	fmt.Println("eq1 = x + (vx * t1) == ", stones[0].x, " + (", stones[0].vx, " * t1)")
	fmt.Println("eq2 = y + (vy * t1) == ", stones[0].y, " + (", stones[0].vy, " * t1)")
	fmt.Println("eq3 = z + (vz * t1) == ", stones[0].z, " + (", stones[0].vz, " * t1)")
	fmt.Println("eq4 = x + (vx * t2) == ", stones[1].x, " + (", stones[1].vx, " * t2)")
	fmt.Println("eq5 = y + (vy * t2) == ", stones[1].y, " + (", stones[1].vy, " * t2)")
	fmt.Println("eq6 = z + (vz * t2) == ", stones[1].z, " + (", stones[1].vz, " * t2)")
	fmt.Println("eq7 = x + (vx * t3) == ", stones[2].x, " + (", stones[2].vx, " * t3)")
	fmt.Println("eq8 = y + (vy * t3) == ", stones[2].y, " + (", stones[2].vy, " * t3)")
	fmt.Println("eq9 = z + (vz * t3) == ", stones[2].z, " + (", stones[2].vz, " * t3)")
	fmt.Println("print(solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9],x,y,z,vx,vy,vz,t1,t2,t3))")
	fmt.Println()

	// Run the above script in SageMath and sum the values of x, y & z
	fmt.Println("Part 2: ", 422521403380479+268293246383898+153073450808511) //843888100572888
}

func parseNums(str string) (int, int, int) {
	parts := strings.Split(str, ", ")
	x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
	return x, y, z
}
