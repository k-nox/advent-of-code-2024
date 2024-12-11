package day04

import (
	"bufio"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/parse"
)

type Point struct {
	X int
	Y int
}

type Grid map[Point]string

func PartOne(useSample bool) int {
	f := parse.OpenInput(2024, 4, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	xLocs := []Point{}
	grid := Grid{}
	y := 0
	for scanner.Scan() {
		curr := strings.TrimSpace(scanner.Text())
		for x := 0; x < len(curr); x++ {
			p := Point{
				X: x,
				Y: y,
			}
			grid[p] = string(curr[x])
			if grid[p] == "X" {
				xLocs = append(xLocs, p)
			}
		}
		y++
	}

	want := []string{"M", "A", "S"}

	count := 0
	for _, xLoc := range xLocs {
		// check going left
		if grid.CheckDirection(Left(xLoc), want, Left) {
			count++
		}
		// check going right
		if grid.CheckDirection(Right(xLoc), want, Right) {
			count++
		}
		// check going up
		if grid.CheckDirection(Up(xLoc), want, Up) {
			count++
		}
		// check going down
		if grid.CheckDirection(Down(xLoc), want, Down) {
			count++
		}
		// check diagonals
		if grid.CheckDirection(UpRight(xLoc), want, UpRight) {
			count++
		}
		if grid.CheckDirection(UpLeft(xLoc), want, UpLeft) {
			count++
		}
		if grid.CheckDirection(DownRight(xLoc), want, DownRight) {
			count++
		}
		if grid.CheckDirection(DownLeft(xLoc), want, DownLeft) {
			count++
		}
	}

	return count
}

func (g Grid) CheckDirection(p Point, want []string, action func(p Point) Point) bool {
	for _, w := range want {
		if l, ok := g[p]; !ok || l != w {
			return false
		}
		p = action(p)
	}
	return true
}

func Left(p Point) Point {
	return Point{X: p.X - 1, Y: p.Y}
}

func Right(p Point) Point {
	return Point{X: p.X + 1, Y: p.Y}
}

func Up(p Point) Point {
	return Point{X: p.X, Y: p.Y + 1}
}

func Down(p Point) Point {
	return Point{X: p.X, Y: p.Y - 1}
}

func UpLeft(p Point) Point {
	return Up(Left(p))
}

func UpRight(p Point) Point {
	return Up(Right(p))
}

func DownLeft(p Point) Point {
	return Down(Left(p))
}

func DownRight(p Point) Point {
	return Down(Right(p))
}
