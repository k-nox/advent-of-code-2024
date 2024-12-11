package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(useSample bool) int {
	re := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)

	file := "input"
	if useSample {
		file = "sample"
	}

	f, err := os.ReadFile(fmt.Sprintf("input/2024/day03/%s.txt", file))
	if err != nil {
		panic(err)
	}

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
