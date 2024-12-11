package day08

import (
	"bufio"

	"github.com/k-nox/advent-of-code-solutions/helper"
	"github.com/k-nox/aoc/util"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2024, 8, useSample)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	grid, antennas := parseInp(scanner)
	uniqueAntinodes := map[util.Point]bool{}

	for _, locations := range antennas {
		for i := 0; i < len(locations)-1; i++ {
			for j := 0; j < len(locations); j++ {
				if i == j {
					continue
				}
				antinodesForPair := findAllAntinodesForPair(grid, locations[i], locations[j])
				for _, n := range antinodesForPair {
					if _, ok := uniqueAntinodes[n]; !ok {
						uniqueAntinodes[n] = true
					}
				}
			}
		}
	}

	return len(uniqueAntinodes)
}

func findAllAntinodesForPair(g util.Grid, a util.Point, b util.Point) []util.Point {
	antinodes := []util.Point{a, b}
	distance := a.Sub(b)
	antinodeA := a.Add(distance)
	_, aOk := g[antinodeA]
	for aOk {
		// keep adding distance to a until it's not valid
		antinodes = append(antinodes, antinodeA)
		antinodeA = antinodeA.Add(distance)
		_, aOk = g[antinodeA]
	}

	antinodeB := b.Sub(distance)
	_, bOk := g[antinodeB]
	for bOk {
		antinodes = append(antinodes, antinodeB)
		antinodeB = antinodeB.Sub(distance)
		_, bOk = g[antinodeB]
	}

	return antinodes
}
