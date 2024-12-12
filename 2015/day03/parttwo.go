package day03

import (
	"image"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	inp := string(helper.ReadInput(2015, 3, useSample))

	santa := image.Point{}
	roboSanta := image.Point{}
	isRobo := false
	visited := helper.NewSet(santa)

	directions := map[rune]helper.Direction{
		'^': helper.Up,
		'>': helper.Right,
		'v': helper.Down,
		'<': helper.Left,
	}

	for _, dir := range inp {
		if isRobo {
			roboSanta = directions[dir](roboSanta)
			visited.Add(roboSanta)
		} else {
			santa = directions[dir](santa)
			visited.Add(santa)
		}
		isRobo = !isRobo
	}

	return visited.Len()
}
