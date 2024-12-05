package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/k-nox/advent-of-code-2024/cli"
	"github.com/k-nox/advent-of-code-2024/util"
)

func main() {
	var useSample bool
	var dayNum int
	var gen bool

	today := time.Now().Day()

	flag.BoolVar(&useSample, "sample", false, "use the sample input")
	flag.IntVar(&dayNum, "day", today, "day to run")
	flag.BoolVar(&gen, "gen", false, "generate files")
	flag.Parse()

	day, ok := cli.GetDay(dayNum)
	if !ok && !gen {
		fmt.Printf("unregistered day requested: %d\n", dayNum)
		os.Exit(1)
	}

	if gen {
		err := util.Generate(dayNum)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("generated files")
		os.Exit(0)
	}

	fmt.Printf("solution for day %d part one: %d\n", dayNum, day.PartOne(useSample))
	fmt.Printf("solution for day %d part two: %d\n", dayNum, day.PartTwo(useSample))
}
