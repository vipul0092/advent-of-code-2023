package day2

import (
	"advent-of-code/reader"
	"fmt"
	"strconv"
	"strings"
)

const input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

const red = 12
const green = 13
const blue = 14

func Solve() {
	//lines := strings.Split(input, "\n")
	lines := strings.Split(reader.Read("/day2/input"), "\n")

	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])

		rounds := strings.Split(parts[1], "; ")
		valid := true
		for _, round := range rounds {
			r := 0
			g := 0
			b := 0
			balls := strings.Split(round, ", ")

			for _, ball := range balls {
				colors := strings.Split(ball, " ")
				cnt, _ := strconv.Atoi(colors[0])
				switch colors[1] {
				case "blue":
					b = cnt
				case "red":
					r = cnt
				case "green":
					g = cnt
				}

				if r > red || g > green || b > blue {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			sum += id
		}
	}
	fmt.Println("Part 1: ", sum)

	// Part 2 Begins
	sum2 := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rounds := strings.Split(parts[1], "; ")
		r := 0
		g := 0
		b := 0
		for _, round := range rounds {
			balls := strings.Split(round, ", ")
			for _, ball := range balls {
				colors := strings.Split(ball, " ")
				cnt, _ := strconv.Atoi(colors[0])
				switch colors[1] {
				case "blue":
					b = max(b, cnt)
				case "red":
					r = max(r, cnt)
				case "green":
					g = max(g, cnt)
				}
			}
		}
		sum2 += r * g * b
	}

	fmt.Println("Part 2: ", sum2)
}
