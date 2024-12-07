// Code generated; DO NOT EDIT.
// This file was generated at
// 2024-12-07 09:22:14.586643 -0800 PST m=+0.003768668
package main

import (
	"log"
	"os"
	"github.com/k-nox/aoc/cli"
	"github.com/k-nox/advent-of-code-2024/day01"
	"github.com/k-nox/advent-of-code-2024/day02"
	"github.com/k-nox/advent-of-code-2024/day03"
	"github.com/k-nox/advent-of-code-2024/day04"
	"github.com/k-nox/advent-of-code-2024/day05"
	"github.com/k-nox/advent-of-code-2024/day06"
	"github.com/k-nox/advent-of-code-2024/day07"
)

var registry = cli.Registry{
	1: {PartOne: day01.PartOne, PartTwo: day01.PartTwo},
	2: {PartOne: day02.PartOne, PartTwo: day02.PartTwo},
	3: {PartOne: day03.PartOne, PartTwo: day03.PartTwo},
	4: {PartOne: day04.PartOne, PartTwo: day04.PartTwo},
	5: {PartOne: day05.PartOne, PartTwo: day05.PartTwo},
	6: {PartOne: day06.PartOne, PartTwo: day06.PartTwo},
	7: {PartOne: day07.PartOne, PartTwo: day07.PartTwo},
}	

func main() {
	app := cli.App(registry, "github.com/k-nox/advent-of-code-2024")
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
