package day07

import (
	"strconv"
	"strings"

	"github.com/k-nox/aoc/util"
)

func PartOne(useSample bool) int {
	f := util.NewScannerForInput(2024, 7, useSample)
	defer f.Close()

	sum := 0
	for f.Scan() {
		target, vals := parse(f.Text())
		if isPossible(target, vals) {
			sum += target
		}
	}

	return sum
}

func parse(line string) (int, []int) {
	var (
		target int
		vals   []int
		err    error
	)
	parts := strings.Split(line, ":")

	target, err = strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	rawVals := strings.Fields(parts[1])
	for _, rawVal := range rawVals {
		val, err := strconv.Atoi(rawVal)
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
	}

	return target, vals
}

func isPossible(target int, vals []int) bool {
	tail := vals[len(vals)-1]
	if len(vals) == 1 {
		return target == tail
	}

	if target%tail == 0 && isPossible(target/tail, vals[:len(vals)-1]) {
		return true
	}
	return isPossible(target-tail, vals[:len(vals)-1])
}
