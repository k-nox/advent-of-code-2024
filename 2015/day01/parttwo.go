package day01

import (
	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	inp := helper.ReadInput(2015, 1, useSample)

	floor := 0
	for pos, char := range inp {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor < 0 {
			return pos + 1
		}
	}

	return 0
}
