package day06

import (
	"bufio"
	"image"
	"slices"

	"github.com/k-nox/advent-of-code-solutions/helper"
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

var dirs = []func(p image.Point) image.Point{
	helper.Up,
	helper.Right,
	helper.Down,
	helper.Left,
}

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2024, 6, useSample)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	grid, guardStart := parseInp(scanner)
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

func isLoop(g grid, d direction, guard image.Point) bool {
	freq := map[image.Point]int{}

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

func visitedCells(g grid, d direction, guard image.Point) []image.Point {
	var visited []image.Point
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

func nextPoint(g grid, d direction, guard image.Point, rotatedN int) (image.Point, direction) {
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
