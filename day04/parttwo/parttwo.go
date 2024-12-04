package parttwo

import (
	"strings"

	"github.com/k-nox/advent-of-code-2024/day04/partone"
	"github.com/k-nox/advent-of-code-2024/util"
)

func Solve() int {
	f := util.NewScannerForInput(4, false)
	defer f.Close()

	grid := partone.Grid{}
	y := 0
	aLocs := []partone.Point{}
	for f.Scan() {
		curr := strings.TrimSpace(f.Text())
		for x := 0; x < len(curr); x++ {
			p := partone.Point{
				X: x,
				Y: y,
			}
			l := string(curr[x])
			grid[p] = l
			if l == "A" {
				aLocs = append(aLocs, p)
			}
		}
		y++
	}

	count := 0
	for _, aLoc := range aLocs {
		topLeft, ok := checkAndGetLetter(grid, partone.UpLeft(aLoc))
		if !ok {
			continue
		}
		topRight, ok := checkAndGetLetter(grid, partone.UpRight(aLoc))
		if !ok {
			continue
		}
		bottomRight, ok := checkAndGetLetter(grid, partone.DownRight(aLoc))
		if !ok {
			continue
		}
		bottomLeft, ok := checkAndGetLetter(grid, partone.DownLeft(aLoc))
		if !ok {
			continue
		}

		validDiagonals := 0
		switch topLeft {
		case "M":
			if bottomRight == "S" {
				validDiagonals++
			}
		case "S":
			if bottomRight == "M" {
				validDiagonals++
			}
		}

		switch topRight {
		case "M":
			if bottomLeft == "S" {
				validDiagonals++
			}
		case "S":
			if bottomLeft == "M" {
				validDiagonals++
			}
		}

		if validDiagonals == 2 {
			count++
		}
	}
	return count
}

func checkAndGetLetter(g partone.Grid, p partone.Point) (string, bool) {
	l, ok := g[p]
	if !ok || (l != "M" && l != "S") {
		return "", false
	}

	return l, ok
}
