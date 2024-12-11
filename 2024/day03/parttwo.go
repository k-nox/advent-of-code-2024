package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	re := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)|(do\(\))|(don't\(\))`)

	file := "input.txt"
	if useSample {
		file = "sample2.txt"
	}

	f := helper.ReadFile(2024, 3, file)

	matches := re.FindAllStringSubmatch(string(f), -1)

	total := 0
	enabled := true
	for _, match := range matches {
		label := match[0]
		if label == "do()" {
			enabled = true
			continue
		}

		if label == "don't()" {
			enabled = false
			continue
		}

		if enabled {
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
	}

	return total
}
