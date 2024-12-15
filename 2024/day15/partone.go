package day15

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"slices"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

const (
	wall  = '#'
	box   = 'O'
	robot = '@'
	empty = '.'
	up    = '^'
	down  = 'v'
	left  = '<'
	right = '>'
)

type warehouse struct {
	walls  helper.Set[image.Point]
	boxes  helper.Set[image.Point]
	bounds image.Rectangle
	robot  image.Point
}

func PartOne(useSample bool) int {
	f := helper.OpenInput(2024, 15, useSample)
	defer f.Close()

	wh, directions := parseInp(f)
	for _, direction := range directions {
		wh.move(direction)
	}

	return wh.gpsSum()
}

func parseInp(f *os.File) (*warehouse, []helper.Direction) {
	var (
		gridInp    []string
		directions []helper.Direction
	)
	scanner := bufio.NewScanner(f)

	parsingDirections := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			parsingDirections = true
			continue
		}

		if parsingDirections {
			directions = slices.Concat(directions, parseDirections(line))
			continue
		}

		gridInp = append(gridInp, line)
	}

	return parseGrid(gridInp), directions
}

func parseDirections(line string) []helper.Direction {
	directions := []helper.Direction{}
	for _, char := range line {
		switch char {
		case up:
			directions = append(directions, helper.Up)
		case down:
			directions = append(directions, helper.Down)
		case left:
			directions = append(directions, helper.Left)
		case right:
			directions = append(directions, helper.Right)
		}
	}

	return directions
}

func parseGrid(inp []string) *warehouse {
	w := &warehouse{
		walls:  helper.NewSet[image.Point](),
		boxes:  helper.NewSet[image.Point](),
		bounds: image.Rect(0, 0, len(inp[0]), len(inp)),
	}

	for y, row := range inp {
		for x, char := range row {
			if char == wall {
				w.walls.Add(image.Pt(x, y))
			}
			if char == robot {
				w.robot = image.Pt(x, y)
			}
			if char == box {
				w.boxes.Add(image.Pt(x, y))
			}
		}
	}

	return w
}

func (w *warehouse) String() string {
	var b strings.Builder
	for y := w.bounds.Min.Y; y < w.bounds.Max.Y; y++ {
		for x := w.bounds.Min.X; x < w.bounds.Max.X; x++ {
			toPrint := empty
			if w.walls.Contains(image.Pt(x, y)) {
				toPrint = wall
			} else if w.boxes.Contains(image.Pt(x, y)) {
				toPrint = box
			} else if image.Pt(x, y) == w.robot {
				toPrint = robot
			}
			fmt.Fprint(&b, string(toPrint))
		}
		fmt.Fprint(&b, "\n")
	}
	return b.String()
}

func (w *warehouse) move(direction helper.Direction) {
	nextPoint := direction(w.robot)
	if w.walls.Contains(nextPoint) {
		return
	}

	canMove := true
	if w.boxes.Contains(nextPoint) {
		canMove = w.moveBox(nextPoint, direction)
	}

	if canMove {
		w.robot = nextPoint
	}
}

func (w *warehouse) moveBox(box image.Point, direction helper.Direction) bool {
	nextPoint := direction(box)
	if w.walls.Contains(nextPoint) {
		return false
	}

	canMove := true
	if w.boxes.Contains(nextPoint) {
		canMove = w.moveBox(nextPoint, direction)
	}

	if canMove {
		w.boxes.Remove(box)
		w.boxes.Add(nextPoint)
	}

	return canMove
}

func (w *warehouse) gpsSum() int {
	sum := 0
	for box := range w.boxes {
		sum += 100*box.Y + box.X
	}
	return sum
}
