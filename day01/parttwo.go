package day01

import (
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-2024/util"
)

func PartTwo() int {
	f := util.NewScannerForInput(1, false)
	defer f.Close()

	var left []int
	rightFreq := make(map[int]int)

	for f.Scan() {
		curr := f.Text()
		nums := strings.Fields(curr)

		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftNum)
		rightFreq[rightNum]++
	}

	score := 0
	for _, num := range left {
		score += (num * rightFreq[num])
	}

	return score
}
