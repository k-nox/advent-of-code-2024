package day01

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/parse"
)

func PartTwo(useSample bool) int {
	f := parse.OpenInput(2024, 1, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var left []int
	rightFreq := make(map[int]int)

	for scanner.Scan() {
		curr := scanner.Text()
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
