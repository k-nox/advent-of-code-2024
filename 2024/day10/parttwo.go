package day10

import (
	"bufio"

	"github.com/k-nox/advent-of-code-solutions/helper"
	"github.com/k-nox/aoc/util"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2024, 10, useSample)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	g, trailheads := parseInp(scanner)
	sum := 0
	for _, trailhead := range trailheads {
		sum += score(g, trailhead, map[util.Point]bool{}, false)
	}

	return sum
}
