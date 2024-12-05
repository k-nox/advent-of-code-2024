// Code generated; DO NOT EDIT.
// This file was generated at
// 2024-12-05 15:02:05.011051 -0800 PST m=+0.003849001
package cli

import (
	"github.com/k-nox/advent-of-code-2024/day01"
	"github.com/k-nox/advent-of-code-2024/day02"
	"github.com/k-nox/advent-of-code-2024/day03"
	"github.com/k-nox/advent-of-code-2024/day04"
	"github.com/k-nox/advent-of-code-2024/day05"
)

type part func(useSample bool) int
type day struct {
	PartOne part
	PartTwo part
}

var registry = map[int]day{
	1: {day01.PartOne, day01.PartTwo},
	2: {day02.PartOne, day02.PartTwo},
	3: {day03.PartOne, day03.PartTwo},
	4: {day04.PartOne, day04.PartTwo},
	5: {day05.PartOne, day05.PartTwo},
}
