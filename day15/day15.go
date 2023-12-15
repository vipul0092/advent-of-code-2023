package day15

import (
	"advent-of-code/utils"
	"fmt"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"strconv"
	"strings"
)

func Solve() {
	items := strings.Split(utils.Read("/day15/input"), ",")

	part1 := 0
	boxes := make(map[int]*linkedhashmap.Map)

	for _, item := range items {
		part1 += getHash(item)

		focallength, label := -1, ""
		if strings.ContainsRune(item, '=') {
			lbl, fstr, _ := strings.Cut(item, "=")
			fl, _ := strconv.Atoi(fstr)
			focallength, label = fl, lbl
		} else {
			lbl, _, _ := strings.Cut(item, "-")
			label = lbl
		}
		box := getHash(label)
		_, has := boxes[box]
		if !has {
			boxes[box] = linkedhashmap.New()
		}
		if focallength == -1 {
			boxes[box].Remove(label)
		} else {
			boxes[box].Put(label, focallength)
		}
	}

	part2 := 0
	for i := 0; i < 256; i++ {
		lensmap, has := boxes[i]
		if !has {
			continue
		}
		for idx, focallength := range lensmap.Values() {
			part2 += (i + 1) * (idx + 1) * focallength.(int)
		}
	}

	fmt.Println("Part 1: ", part1) // 510792
	fmt.Println("Part 2: ", part2) // 269410
}

func getHash(str string) int {
	hash := 0
	for _, char := range str {
		if char == '\n' {
			continue
		}
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return hash
}
