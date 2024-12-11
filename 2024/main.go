// Code generated; DO NOT EDIT.
// This file was generated at
// 2024-12-11 09:25:37.963079 -0800 PST m=+0.002737418
package main

import (
	"log"
	"os"
	"github.com/k-nox/aoc/cli"
	"github.com/k-nox/advent-of-code-solutions/2024/day01"
	"github.com/k-nox/advent-of-code-solutions/2024/day02"
	"github.com/k-nox/advent-of-code-solutions/2024/day03"
	"github.com/k-nox/advent-of-code-solutions/2024/day04"
	"github.com/k-nox/advent-of-code-solutions/2024/day05"
	"github.com/k-nox/advent-of-code-solutions/2024/day06"
	"github.com/k-nox/advent-of-code-solutions/2024/day07"
	"github.com/k-nox/advent-of-code-solutions/2024/day08"
	"github.com/k-nox/advent-of-code-solutions/2024/day09"
	"github.com/k-nox/advent-of-code-solutions/2024/day10"
	"github.com/k-nox/advent-of-code-solutions/2024/day11"
)

var registry = cli.Registry{
	"day01": {PartOne: day01.PartOne, PartTwo: day01.PartTwo},
	"day02": {PartOne: day02.PartOne, PartTwo: day02.PartTwo},
	"day03": {PartOne: day03.PartOne, PartTwo: day03.PartTwo},
	"day04": {PartOne: day04.PartOne, PartTwo: day04.PartTwo},
	"day05": {PartOne: day05.PartOne, PartTwo: day05.PartTwo},
	"day06": {PartOne: day06.PartOne, PartTwo: day06.PartTwo},
	"day07": {PartOne: day07.PartOne, PartTwo: day07.PartTwo},
	"day08": {PartOne: day08.PartOne, PartTwo: day08.PartTwo},
	"day09": {PartOne: day09.PartOne, PartTwo: day09.PartTwo},
	"day10": {PartOne: day10.PartOne, PartTwo: day10.PartTwo},
	"day11": {PartOne: day11.PartOne, PartTwo: day11.PartTwo},
}	

func main() {
	app := cli.App(registry)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
