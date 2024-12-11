package day06

import (
	"slices"

	"github.com/k-nox/aoc/util"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

func (d direction) next() direction {
	return (d + 1) % 4
}

var dirs = []func(p util.Point) util.Point{
	util.Up,
	util.Right,
	util.Down,
	util.Left,
}

func PartTwo(useSample bool) int {
	f := util.NewScannerForInput(2024, 6, useSample)
	defer f.Close()
	grid, guardStart := parse(f)
	visited := visitedCells(grid, up, guardStart)

	loops := 0
	for _, cell := range visited {
		grid[cell] = "#"
		if isLoop(grid, up, guardStart) {
			loops++
		}
		grid[cell] = "."
	}

	return loops
}

func isLoop(g util.Grid, d direction, guard util.Point) bool {
	freq := map[util.Point]int{}

	at := guard
	for {
		nextPt, dir := nextPoint(g, d, at, 0)
		if at == nextPt {
			break
		}
		freq[nextPt]++
		if count := freq[nextPt]; count > 4 {
			return true
		}
		d = dir
		at = nextPt
	}
	return false
}

func visitedCells(g util.Grid, d direction, guard util.Point) []util.Point {
	var visited []util.Point
	at := guard
	for {
		nextPt, dir := nextPoint(g, d, at, 0)
		if at == nextPt {
			break
		}
		if !slices.Contains(visited, nextPt) {
			visited = append(visited, nextPt)
		}
		d = dir
		at = nextPt
	}
	return visited
}

func nextPoint(g util.Grid, d direction, guard util.Point, rotatedN int) (util.Point, direction) {
	if rotatedN == 4 {
		return guard, d
	}
	nextPt := dirs[d](guard)
	next, ok := g[nextPt]
	if !ok {
		return guard, d
	}

	if next == "#" {
		return nextPoint(g, d.next(), guard, rotatedN+1)
	}

	return nextPt, d

}
