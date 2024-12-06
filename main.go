// Code generated; DO NOT EDIT.
// This file was generated at
// 2024-12-05 16:55:02.907414 -0800 PST m=+0.005327209
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
)

var registry = cli.Registry{
	1: {PartOne: day01.PartOne, PartTwo: day01.PartTwo},
	2: {PartOne: day02.PartOne, PartTwo: day02.PartTwo},
	3: {PartOne: day03.PartOne, PartTwo: day03.PartTwo},
	4: {PartOne: day04.PartOne, PartTwo: day04.PartTwo},
	5: {PartOne: day05.PartOne, PartTwo: day05.PartTwo},
}	

func main() {
	app := cli.App(registry, "github.com/k-nox/advent-of-code-2024")
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

