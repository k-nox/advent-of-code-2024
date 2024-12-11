package day10

import (
	"github.com/k-nox/aoc/util"
)

func PartTwo(useSample bool) int {
	f := util.NewScannerForInput(2024, 10, useSample)
	defer f.Close()

	g, trailheads := parse(f)
	sum := 0
	for _, trailhead := range trailheads {
		sum += score(g, trailhead, map[util.Point]bool{}, false)
	}

	return sum
}
