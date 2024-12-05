package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/k-nox/advent-of-code-2024/day01"
	"github.com/k-nox/advent-of-code-2024/day02"
	"github.com/k-nox/advent-of-code-2024/day03"
	"github.com/k-nox/advent-of-code-2024/day04"
	"github.com/k-nox/advent-of-code-2024/day05"
)

type part func(useSample bool) int
type day struct {
	partOne part
	partTwo part
}

var registry = map[int]day{
	1: {day01.PartOne, day01.PartTwo},
	2: {day02.PartOne, day02.PartTwo},
	3: {day03.PartOne, day03.PartTwo},
	4: {day04.PartOne, day04.PartTwo},
	5: {day05.PartOne, day05.PartTwo},
}

func main() {
	var useSample bool
	var dayNum int

	today := time.Now().Day()

	flag.BoolVar(&useSample, "sample", false, "use the sample input")
	flag.IntVar(&dayNum, "day", today, "day to run")
	flag.Parse()

	day, ok := registry[dayNum]
	if !ok {
		fmt.Printf("unregistered day requested: %d\n", dayNum)
		os.Exit(1)
	}

	fmt.Printf("solution for day %d part one: %d\n", dayNum, day.partOne(useSample))
	fmt.Printf("solution for day %d part two: %d\n", dayNum, day.partTwo(useSample))
}
