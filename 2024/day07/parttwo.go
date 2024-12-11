package day07

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/parse"
)

func PartTwo(useSample bool) int {
	f := parse.OpenInput(2024, 7, useSample)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		target, vals := parseInp(scanner.Text())
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
