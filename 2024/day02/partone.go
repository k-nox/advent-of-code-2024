package day02

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
	f := helper.OpenInput(2024, 2, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	safe := 0

	for scanner.Scan() {
		raw := scanner.Text()

		report := strings.Fields(raw)

		prev, err := strconv.Atoi(report[0])
		if err != nil {
			panic(err)
		}
		curr, err := strconv.Atoi(report[1])
		if err != nil {
			panic(err)
		}
		increasing := false
		if curr > prev {
			increasing = true
		}

		safeReport := true

		for i := 1; i < len(report); i++ {
			curr, err := strconv.Atoi(report[i])
			if err != nil {
				panic(err)
			}
			// check if correct direction
			if (increasing && prev > curr) || (!increasing && prev < curr) {
				safeReport = false
				break
			}

			if distance := math.Abs(float64(curr) - float64(prev)); distance > 3 || distance < 1 {
				safeReport = false
				break
			}

			prev = curr
		}

		if safeReport {
			safe++
		}
	}
	return safe
}
