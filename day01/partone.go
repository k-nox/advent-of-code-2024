package day01

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-2024/util"
)

func PartOne(useSample bool) int {
	f := util.NewScannerForInput(1, useSample)
	defer f.Close()
	var left []int
	var right []int

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
		right = append(right, rightNum)
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0

	for i, leftNum := range left {
		rightNum := right[i]

		distance += int(math.Abs(float64(leftNum) - float64(rightNum)))
	}

	return distance
}
