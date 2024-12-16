package day15

import (
	"bufio"
	"fmt"
	"image"
	"maps"
	"os"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

const (
	leftBox  = '['
	rightBox = ']'
)

type warehouse2 struct {
	bounds image.Rectangle
	grid   map[image.Point]rune
}

var directions = map[rune]helper.Direction{
	up:    helper.Up,
	left:  helper.Left,
	right: helper.Right,
	down:  helper.Down,
}

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2024, 15, useSample)
	defer f.Close()

	w, bot, dirs := parseInp2(f)
	for _, dir := range dirs {
		grid := w.grid
		bot, grid = move(bot, dir, grid)
		w.grid = grid
	}

	return w.gpsSum()
}

func move(pt image.Point, direction rune, grid map[image.Point]rune) (image.Point, map[image.Point]rune) {
	nextPoint := directions[direction](pt)
	nextChar := grid[nextPoint]
	char := grid[pt]
	gridCopy := maps.Clone(grid)
	if nextChar == empty {
		gridCopy[pt] = empty
		gridCopy[nextPoint] = char
		return nextPoint, gridCopy
	}

	if nextChar == wall {
		return pt, gridCopy
	}

	if direction == left || direction == right {
		movedPt, grid2 := move(nextPoint, direction, gridCopy)
		if movedPt == nextPoint {
			return pt, gridCopy
		}
		gridCopy = grid2
		gridCopy[pt] = empty
		gridCopy[nextPoint] = char
		return nextPoint, gridCopy
	}

	var nextPoint2 image.Point
	if nextChar == leftBox {
		nextPoint2 = helper.Right(nextPoint)
	} else {
		nextPoint2 = helper.Left(nextPoint)
	}

	movedPt, grid2 := move(nextPoint, direction, gridCopy)
	if movedPt != nextPoint {
		movedPt2, grid3 := move(nextPoint2, direction, grid2)
		if movedPt2 != nextPoint2 {
			gridCopy = grid3
			gridCopy[nextPoint] = char
			gridCopy[pt] = empty
			return nextPoint, gridCopy
		}
	}

	return pt, gridCopy

}

func parseInp2(f *os.File) (warehouse2, image.Point, []rune) {
	gridInp := []string{}
	directions := []rune{}

	scanner := bufio.NewScanner(f)
	parsingDirections := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			parsingDirections = true
			continue
		}

		if parsingDirections {
			for _, d := range line {
				directions = append(directions, d)
			}
			continue
		}
		gridInp = append(gridInp, line)
	}

	w, bot := parseWideGrid(gridInp)

	return w, bot, directions
}

func parseWideGrid(inp []string) (warehouse2, image.Point) {
	w := warehouse2{
		bounds: image.Rect(0, 0, len(inp[0])*2, len(inp)),
		grid:   map[image.Point]rune{},
	}
	bot := image.Point{}

	for y, row := range inp {
		x := 0
		for _, char := range row {
			left := image.Pt(x, y)
			right := image.Pt(x+1, y)
			leftChar := char
			rightChar := char
			if char == robot {
				rightChar = empty
				bot = left
			}

			if char == box {
				leftChar = leftBox
				rightChar = rightBox
			}

			w.grid[left] = leftChar
			w.grid[right] = rightChar
			x += 2
		}
	}

	return w, bot
}

func (w warehouse2) String() string {
	var b strings.Builder
	for y := w.bounds.Min.Y; y < w.bounds.Max.Y; y++ {
		for x := w.bounds.Min.X; x < w.bounds.Max.X; x++ {
			fmt.Fprint(&b, string(w.grid[image.Pt(x, y)]))
		}
		fmt.Fprint(&b, "\n")
	}
	return b.String()
}

func (w *warehouse2) gpsSum() int {
	sum := 0
	for pt, char := range w.grid {
		if char == leftBox {
			sum += 100*pt.Y + pt.X
		}
	}
	return sum
}
