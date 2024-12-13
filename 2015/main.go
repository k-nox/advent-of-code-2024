// Code generated; DO NOT EDIT.
// This file was generated at
// 2024-12-13 10:31:54.598871 -0800 PST m=+0.631760793
package main

import (
	"log"
	"os"
	"github.com/k-nox/aoc/cli"
	"github.com/k-nox/advent-of-code-solutions/2015/day01"
	"github.com/k-nox/advent-of-code-solutions/2015/day02"
	"github.com/k-nox/advent-of-code-solutions/2015/day03"
	"github.com/k-nox/advent-of-code-solutions/2015/day04"
	"github.com/k-nox/advent-of-code-solutions/2015/day05"
	"github.com/k-nox/advent-of-code-solutions/2015/day06"
)

var registry = cli.Registry{
	"day01": {PartOne: day01.PartOne, PartTwo: day01.PartTwo},
	"day02": {PartOne: day02.PartOne, PartTwo: day02.PartTwo},
	"day03": {PartOne: day03.PartOne, PartTwo: day03.PartTwo},
	"day04": {PartOne: day04.PartOne, PartTwo: day04.PartTwo},
	"day05": {PartOne: day05.PartOne, PartTwo: day05.PartTwo},
	"day06": {PartOne: day06.PartOne, PartTwo: day06.PartTwo},
}	

func main() {
	app := cli.App(registry)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
