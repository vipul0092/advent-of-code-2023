package day19

import (
	"advent-of-code/utils"
	"fmt"
	"maps"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Rule struct {
	property     rune
	sign         rune
	checkVal     int
	destWorkflow string
}

type Range struct {
	start int
	end   int
}

func Solve() {
	lines := utils.ReadLines("/day19/input")

	workflows := make(map[string][]Rule)
	parts := make([]Part, 0)

	partstart := false
	for _, line := range lines {
		if len(line) == 0 {
			partstart = true
		} else if !partstart {
			populateWorkflow(line, workflows)
		} else {
			line = line[1 : len(line)-1]
			nums := strings.Split(line, ",")
			x, _ := strconv.Atoi(nums[0][2:])
			m, _ := strconv.Atoi(nums[1][2:])
			a, _ := strconv.Atoi(nums[2][2:])
			s, _ := strconv.Atoi(nums[3][2:])
			parts = append(parts, Part{x, m, a, s})
		}
	}

	sum := 0
	for _, part := range parts {
		ranges := make(map[rune]Range)
		ranges['x'] = Range{part.x, part.x}
		ranges['m'] = Range{part.m, part.m}
		ranges['a'] = Range{part.a, part.a}
		ranges['s'] = Range{part.s, part.s}

		ways := evaluateWays(workflows, "in", ranges)
		if ways > 0 {
			sum += part.x + part.m + part.a + part.s
		}
	}
	fmt.Println("Part 1: ", sum) // 330820

	ranges := make(map[rune]Range)
	ranges['x'] = Range{1, 4000}
	ranges['m'] = Range{1, 4000}
	ranges['a'] = Range{1, 4000}
	ranges['s'] = Range{1, 4000}
	fmt.Println("Part 2: ", evaluateWays(workflows, "in", ranges)) // 123972546935551
}

func evaluateWays(workflows map[string][]Rule, workflow string, ranges map[rune]Range) int {
	rules, has := workflows[workflow]
	if !has {
		if workflow == "A" {
			return getAllWays(ranges)
		}
		return 0
	}

	sum := 0
	for _, rule := range rules {
		if rule.checkVal != -1 {
			sign := rule.sign
			prop := rule.property
			checkVal := rule.checkVal
			destWorkflow := rule.destWorkflow
			if sign == '>' {
				r := ranges[prop]
				if r.end <= checkVal {
					continue
				} else if r.start > checkVal {
					sum += evaluateWays(workflows, destWorkflow, ranges)
					break
				} else {
					newRange := make(map[rune]Range)
					maps.Copy(newRange, ranges)
					newRange[prop] = Range{checkVal + 1, ranges[prop].end}
					sum += evaluateWays(workflows, destWorkflow, newRange)

					ranges[prop] = Range{ranges[prop].start, checkVal}
				}
			} else if sign == '<' {
				r := ranges[prop]
				if r.start >= checkVal {
					continue
				} else if r.end < checkVal {
					sum += evaluateWays(workflows, destWorkflow, ranges)
					break
				} else {
					newRange := make(map[rune]Range)
					maps.Copy(newRange, ranges)
					newRange[prop] = Range{ranges[prop].start, checkVal - 1}
					sum += evaluateWays(workflows, destWorkflow, newRange)

					ranges[prop] = Range{checkVal, ranges[prop].end}
				}
			}
		} else {
			sum += evaluateWays(workflows, rule.destWorkflow, ranges)
			break
		}
	}
	return sum
}

func getAllWays(ranges map[rune]Range) int {
	result := 1
	for _, r := range ranges {
		result *= r.end - r.start + 1
	}
	return result
}

func populateWorkflow(line string, workflows map[string][]Rule) {
	var parts = strings.Split(line, "{")
	name := parts[0]
	parts[1] = parts[1][:len(parts[1])-1]
	rules := strings.Split(parts[1], ",")
	workflows[name] = make([]Rule, len(rules))

	for i, rulestr := range rules {
		if strings.ContainsRune(rulestr, '>') || strings.ContainsRune(rulestr, '<') {
			parts = strings.Split(rulestr, ":")
			dest := parts[1]
			sign := rune(parts[0][1])
			prop := rune(parts[0][0])
			value, _ := strconv.Atoi(parts[0][2:])
			workflows[name][i] = Rule{prop, sign, value, dest}
		} else {
			workflows[name][i] = Rule{checkVal: -1, destWorkflow: rulestr}
		}
	}
}
