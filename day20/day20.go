package day20

import (
	"advent-of-code/utils"
	"fmt"
	"github.com/emirpasic/gods/queues/arrayqueue"
	"strings"
)

type Pulse struct {
	from string
	to   string
	high bool
}

var broadcaster []string
var flipFlopStates map[string]bool
var flipFlopDest map[string][]string
var conjunctionStates map[string]map[string]bool
var conjunctionDest map[string][]string
var inputModuleToRx string

func Solve() {
	parseInput(utils.ReadLines("/day20/input"))
	fmt.Println("Part 1: ", runSimulation(1000, false))

	parseInput(utils.ReadLines("/day20/input"))
	fmt.Println("Part 2: ", runSimulation(20000, true))
}

func runSimulation(presses int, calcCycle bool) int {
	highcnt, lowcnt, cycles := 0, 0, make(map[string][]int)

	for i := 0; i < presses; i++ {
		queue := arrayqueue.New()
		queue.Enqueue(Pulse{"", "broadcaster", false})

		for !queue.Empty() {
			p, _ := queue.Dequeue()
			pulse := p.(Pulse)
			from, module, high := pulse.from, pulse.to, pulse.high

			if high {
				highcnt++
			} else {
				lowcnt++
			}

			if calcCycle && high && module == inputModuleToRx {
				if _, has := cycles[from]; !has {
					cycles[from] = make([]int, 0)
				}
				cycles[from] = append(cycles[from], i+1)
			}

			if module == "broadcaster" {
				for _, d := range broadcaster {
					queue.Enqueue(Pulse{module, d, high})
				}
			} else if current, hasF := flipFlopStates[module]; hasF && !high { // flip flops ignore high pulse
				flipFlopStates[module] = !current
				for _, d := range flipFlopDest[module] {
					queue.Enqueue(Pulse{module, d, !current})
				}
			} else if _, hasC := conjunctionStates[module]; hasC {
				conjunctionStates[module][from] = high
				allHigh := true
				for _, h := range conjunctionStates[module] {
					allHigh = allHigh && h
				}
				for _, d := range conjunctionDest[module] {
					queue.Enqueue(Pulse{module, d, !allHigh})
				}
			}
		}
	}

	if !calcCycle {
		return lowcnt * highcnt
	} else {
		numbers := make([]int, 0)
		for _, list := range cycles {
			numbers = append(numbers, list[1]-list[0])
		}
		return utils.LCM(numbers)
	}
}

func parseInput(lines []string) {
	broadcaster, flipFlopStates, flipFlopDest, conjunctionStates, conjunctionDest =
		make([]string, 0), make(map[string]bool), make(map[string][]string),
		make(map[string]map[string]bool), make(map[string][]string)

	for _, line := range lines {
		if strings.HasPrefix(line, "broadcaster") {
			broadcaster = append(broadcaster, strings.Split(strings.Split(line, " -> ")[1], ", ")...)
		} else if strings.ContainsRune(line, '%') {
			module, dest, _ := strings.Cut(line, " -> ")
			module = module[1:]
			flipFlopDest[module] = strings.Split(dest, ", ")
			flipFlopStates[module] = false
		} else if strings.ContainsRune(line, '&') {
			module, dest, _ := strings.Cut(line, " -> ")
			module = module[1:]
			conjunctionDest[module] = strings.Split(dest, ", ")
			conjunctionStates[module] = make(map[string]bool)
			for _, d := range conjunctionDest[module] {
				if d == "rx" {
					inputModuleToRx = module
				}
			}
		}
	}

	for ff, ffd := range flipFlopDest {
		for _, d := range ffd {
			if _, has := conjunctionStates[d]; has {
				conjunctionStates[d][ff] = false
			}
		}
	}
}
