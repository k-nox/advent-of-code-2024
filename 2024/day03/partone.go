package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
	re := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)

	f := helper.ReadInput(2024, 3, useSample)

	matches := re.FindAllStringSubmatch(string(f), -1)

	total := 0

	for _, match := range matches {
		nums := strings.Split(match[1], ",")
		a, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		total += (a * b)
	}

	return total
}
