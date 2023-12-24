package day24

import (
	"advent-of-code/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
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

type Point struct {
	x int
	y int
	z int
}

func Solve() {
	lines := utils.ReadLines("/day24/input")

	stones := make([]Stone, len(lines))

	vxx, vyy, vzz := make(map[int][]int), make(map[int][]int), make(map[int][]int)
	for i, line := range lines {
		p, v, _ := strings.Cut(line, " @ ")
		x, y, z := parseNums(p)
		vx, vy, vz := parseNums(v)

		m := float64(vy) / float64(vx)     // y2-y1 / x2-x1
		c := float64(y) - (m * float64(x)) // y = mx + c => c = y - mx
		// y = mx + c => mx - y + c = 0
		stones[i] = Stone{x, y, z, vx, vy, vz, Line{m, -1, c}}
		if _, hx := vxx[vx]; !hx {
			vxx[vx] = make([]int, 0)
		}
		if _, hy := vyy[vy]; !hy {
			vyy[vy] = make([]int, 0)
		}
		if _, hz := vzz[vz]; !hz {
			vzz[vz] = make([]int, 0)
		}
		vxx[vx] = append(vxx[vx], x)
		vyy[vy] = append(vyy[vy], y)
		vzz[vz] = append(vzz[vz], z)
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

	// ------ Part 2 Begins ------
	// Algebraic approach inspired by: https://www.reddit.com/r/adventofcode/comments/18pnycy/2023_day_24_solutions/keqf8uq/
	// If two hailstones have the same velocity in a direction, then the distance between them in that direction
	// must be a multiple of the Rock velocity - hailstone velocity
	// We iterate on Rock velocities on from -1000 to 1000, and find the valid velocities in each direction
	// Interestingly we only get one correct velocity in each direction!
	rvx, rvy, rvz := getPossibleVelocity(vxx), getPossibleVelocity(vyy), getPossibleVelocity(vzz)
	fmt.Println(rvx, rvy, rvz)

	vx1, vy1, vz1 := stones[0].vx-rvx, stones[0].vy-rvy, stones[0].vz-rvz
	vx2, vy2, vz2 := stones[1].vx-rvx, stones[1].vy-rvy, stones[1].vz-rvz

	// Intersection between two 3D lines logic taken from: https://paulbourke.net/geometry/pointlineplane/
	p1 := Point{stones[0].x + vx1, stones[0].y + vy1, stones[0].z + vz1}
	p2 := Point{p1.x + vx1, p1.y + vy1, p1.z + vz1}
	p3 := Point{stones[1].x + vx2, stones[1].y + vy2, stones[1].z + vz2}
	p4 := Point{p3.x + vx2, p3.y + vy2, p3.z + vz2}

	// Pa = P1 + mua (P2 - P1)
	// mua = ( d1343 d4321 - d1321 d4343 ) / ( d2121 d4343 - d4321 d4321 )
	// where dmnop = (xm - xn)(xo - xp) + (ym - yn)(yo - yp) + (zm - zn)(zo - zp)
	mua := ((dmnop(p1, p3, p4, p3) * dmnop(p4, p3, p2, p1)) - (dmnop(p1, p3, p2, p1) * dmnop(p4, p3, p4, p3))) /
		((dmnop(p2, p1, p2, p1) * dmnop(p4, p3, p4, p3)) - (dmnop(p4, p3, p2, p1) * dmnop(p4, p3, p2, p1)))

	px := float64(p1.x) + (mua * float64(p2.x-p1.x))
	py := float64(p1.y) + (mua * float64(p2.y-p1.y))
	pz := float64(p1.z) + (mua * float64(p2.z-p1.z))
	fmt.Println(int(px + py + pz)) // 843888100572888

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

// dmnop = (xm - xn)(xo - xp) + (ym - yn)(yo - yp) + (zm - zn)(zo - zp)
func dmnop(m, n, o, p Point) float64 {
	return float64(((m.x - n.x) * (o.x - p.x)) + ((m.y - n.y) * (o.y - p.y)) + ((m.z - n.z) * (o.z - p.z)))
}

func getPossibleVelocity(vx map[int][]int) int {
	possvx := mapset.NewSet[int]()
	for v, d := range vx {
		if len(d) == 1 {
			continue
		}
		possible := mapset.NewSet[int]()
		for i := 0; i < len(d)-1; i++ {
			for j := i + 1; j < len(d); j++ {
				ddiff := abs(d[i] - d[j])
				for pv := -1000; pv <= 1000; pv++ {
					if pv != v && ddiff%(pv-v) == 0 {
						possible.Add(pv)
					}
				}
			}
		}
		if possvx.IsEmpty() {
			possvx = possible
		} else {
			possvx = possible.Intersect(possvx)
		}
	}

	if possvx.Cardinality() != 1 {
		panic("Not possible to solve!")
	}
	val, _ := possvx.Pop()
	return val
}

func parseNums(str string) (int, int, int) {
	parts := strings.Split(str, ", ")
	x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
	return x, y, z
}

func abs(d int) int {
	if d < 0 {
		return -d
	}
	return d
}
