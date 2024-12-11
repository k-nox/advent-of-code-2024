package day10

import (
	"bufio"
	"image"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

type grid map[image.Point]int

func PartOne(useSample bool) int {
	f := helper.OpenInput(2024, 10, useSample)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	g, trailheads := parseInp(scanner)
	sum := 0
	for _, trailhead := range trailheads {
		sum += score(g, trailhead, map[image.Point]bool{}, true)
	}
	return sum
}

// returns a grid & a list of trailheads
func parseInp(scanner *bufio.Scanner) (grid, []image.Point) {
	g := grid{}
	trailheads := []image.Point{}

	y := 0
	for scanner.Scan() {
		row := scanner.Text()
		for x, char := range row {
			val := int(char - '0')
			point := image.Point{
				X: x,
				Y: y,
			}
			g[point] = val
			if val == 0 {
				trailheads = append(trailheads, point)
			}
		}
		y++
	}

	return g, trailheads
}

func score(g grid, trailhead image.Point, foundPeaks map[image.Point]bool, mustBeUnique bool) int {
	curr := g[trailhead]

	if curr == 9 {
		if _, seenBefore := foundPeaks[trailhead]; mustBeUnique && seenBefore {
			return 0
		}
		foundPeaks[trailhead] = true
		return 1
	}

	sum := 0

	directions := []func(image.Point) image.Point{
		helper.Up,
		helper.Right,
		helper.Down,
		helper.Left,
	}

	for _, direction := range directions {
		nextPoint := direction(trailhead)
		if val, ok := g[nextPoint]; ok && val == curr+1 {
			sum += score(g, nextPoint, foundPeaks, mustBeUnique)
		}
	}

	return sum
}
