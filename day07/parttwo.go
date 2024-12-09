package day07

import (
	"strconv"
	"strings"

	"github.com/k-nox/aoc/util"
)

func PartTwo(useSample bool) int {
	f := util.NewScannerForInput(7, useSample)
	defer f.Close()

	sum := 0
	for f.Scan() {
		target, vals := parse(f.Text())
		if isPossibleWithConcat(target, vals) {
			sum += target
		}

	}

	return sum
}

func isPossibleWithConcat(target int, vals []int) bool {
	if target < 0 {
		return false
	}
	tail := vals[len(vals)-1]
	if len(vals) == 1 {
		return target == tail
	}

	if target%tail == 0 && isPossibleWithConcat(target/tail, vals[:len(vals)-1]) {
		return true
	}

	targetStr := strconv.Itoa(target)
	tailStr := strconv.Itoa(tail)
	if len(targetStr) > len(tailStr) {
		newTargetStr, endsWithTail := strings.CutSuffix(strconv.Itoa(target), strconv.Itoa(tail))
		newTarget, err := strconv.Atoi(newTargetStr)
		if err != nil {
			panic(err)
		}
		if endsWithTail && isPossibleWithConcat(newTarget, vals[:len(vals)-1]) {
			return true
		}
	}

	return isPossibleWithConcat(target-tail, vals[:len(vals)-1])
}
