package day11

import (
	"bufio"

	"github.com/k-nox/advent-of-code-solutions/parse"
)

func PartTwo(useSample bool) int {
	f := parse.OpenInput(2024, 11, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
	
	}

	return 0
}
