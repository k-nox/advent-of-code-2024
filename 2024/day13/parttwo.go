package day13

import (
	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2024, 13, useSample)
	defer f.Close()

	machines := parseInp(f, true)
	tokens := 0
	for _, mach := range machines {
		toWin, possible := mach.play()
		if possible {
			tokens += toWin
		}
	}
	return tokens
}
