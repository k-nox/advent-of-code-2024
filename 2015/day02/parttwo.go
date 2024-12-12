package day02

import (
	"bufio"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2015, 2, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
	
	}

	return 0
}
