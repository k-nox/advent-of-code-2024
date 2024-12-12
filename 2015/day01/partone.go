package day01

import (
	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
	inp := helper.ReadInput(2015, 1, useSample)

	floor := 0
	for _, char := range inp {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}

	}

	return floor
}
