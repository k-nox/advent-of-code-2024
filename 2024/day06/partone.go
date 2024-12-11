package day06

import (
	"bufio"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
	"github.com/k-nox/aoc/util"
)

func PartOne(useSample bool) int {
	f := helper.OpenInput(2024, 6, useSample)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	grid, guard := parseInp(scanner)

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

func parseInp(scanner *bufio.Scanner) (util.Grid, util.Point) {
	grid := util.Grid{}
	y := 0
	guard := util.Point{}
	for scanner.Scan() {
		curr := strings.TrimSpace(scanner.Text())
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
	return grid, guard
}
