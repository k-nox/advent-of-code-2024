package day02

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/k-nox/aoc/util"
)

func PartTwo(useSample bool) int {
	f := util.NewScannerForInput(2024, 2, useSample)
	defer f.Close()
	safe := 0
	// threshold := 20

	for f.Scan() {
		curr := f.Text()
		rawReport := strings.Fields(curr)
		report := make([]int, len(rawReport))
		for i, s := range rawReport {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			report[i] = num
		}

		if isSafeReport(report, false) {
			safe++
		}

	}

	return safe
}

func isSafeReport(report []int, toleranceUsed bool) bool {
	left := 0
	right := 1
	isIncreasing := false
	if report[right] > report[left] {
		isIncreasing = true
	}

	isSafe := true

	for right < len(report) {
		if isSafePair(report[left], report[right], isIncreasing) {
			left++
			right++
		} else {
			isSafe = false
			break
		}
	}

	if !isSafe && !toleranceUsed {
		// if right is the last index, we can just remove that to be safe
		if right == len(report)-1 {
			isSafe = true
			// check if removing the first element in the makes the report safe
		} else if isSafeReport(report[1:], true) {
			isSafe = true
			// check if removing the left element makes it safe
		} else if isSafeReport(slices.Delete(slices.Clone(report), left, left+1), true) {
			isSafe = true
			// check if removing the right element makes it safe
		} else if isSafeReport(slices.Delete(slices.Clone(report), right, right+1), true) {
			isSafe = true
		}
	}

	return isSafe
}

func isSafePair(prev int, curr int, increasing bool) bool {
	if increasing && prev >= curr {
		return false
	}

	if !increasing && prev <= curr {
		return false
	}

	difference := math.Abs(float64(curr) - float64(prev))
	if difference > 3 || difference < 1 {
		return false
	}

	return true
}
