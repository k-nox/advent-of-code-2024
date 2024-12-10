package day10

import (
	"github.com/k-nox/aoc/util"
)

type grid map[util.Point]int

func PartOne(useSample bool) int {
	f := util.NewScannerForInput(10, useSample)
	defer f.Close()

	g, trailheads := parse(f)
	sum := 0
	for _, trailhead := range trailheads {
		sum += score(g, trailhead, map[util.Point]bool{}, true)
	}
	return sum
}

// returns a grid & a list of trailheads
func parse(f *util.FileScanner) (grid, []util.Point) {
	g := grid{}
	trailheads := []util.Point{}

	y := 0
	for f.Scan() {
		row := f.Text()
		for x, char := range row {
			val := int(char - '0')
			point := util.Point{
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

func score(g grid, trailhead util.Point, foundPeaks map[util.Point]bool, mustBeUnique bool) int {
	curr := g[trailhead]

	if curr == 9 {
		if _, seenBefore := foundPeaks[trailhead]; mustBeUnique && seenBefore {
			return 0
		}
		foundPeaks[trailhead] = true
		return 1
	}

	sum := 0

	directions := []func(util.Point) util.Point{
		util.Up,
		util.Right,
		util.Down,
		util.Left,
	}

	for _, direction := range directions {
		nextPoint := direction(trailhead)
		if val, ok := g[nextPoint]; ok && val == curr+1 {
			sum += score(g, nextPoint, foundPeaks, mustBeUnique)
		}
	}

	return sum
}
