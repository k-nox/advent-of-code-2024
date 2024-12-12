package day03

import (
	"image"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
	inp := string(helper.ReadInput(2015, 3, useSample))

	santa := image.Point{}
	visited := helper.NewSet(santa)
	directions := map[rune]helper.Direction{
		'^': helper.Up,
		'>': helper.Right,
		'v': helper.Down,
		'<': helper.Left,
	}

	for _, dir := range inp {
		santa = directions[dir](santa)
		visited.Add(santa)
	}

	return visited.Len()
}
