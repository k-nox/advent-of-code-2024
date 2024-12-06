package day06

import (
	"strings"

	"github.com/k-nox/aoc/util"
)

func PartOne(useSample bool) int {
	f := util.NewScannerForInput(6, useSample)
	defer f.Close()

	grid := util.Grid{}
	y := 0
	guard := util.Point{}
	for f.Scan() {
		curr := strings.TrimSpace(f.Text())
		for x := 0; x < len(curr); x++ {
			p := util.Point{
				X: x,
				Y: y,
			}
			c := string(curr[x])
			if c == "^" {
				c = "."
				guard = p
			}
			grid[p] = c
		}
		y++
	}

	dirs := []func(p util.Point) util.Point{
		util.Up,
		util.Right,
		util.Down,
		util.Left,
	}

	currDir := 0 // index for Up
	sum := 1
	visited := map[util.Point]bool{
		guard: true,
	}

	for {
		nextPt := dirs[currDir](guard)
		next, ok := grid[nextPt]
		if !ok {
			break
		}

		if next == "#" {
			currDir = (currDir + 1) % 4
			nextPt = dirs[currDir](guard)
			_, ok = grid[nextPt]
			if !ok {
				break
			}
		}
		guard = nextPt
		if _, ok := visited[guard]; !ok {
			sum++
		}
		visited[guard] = true
	}

	return sum
}
