package day02

import (
	"math"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-2024/util"
)

func PartOne() int {
	f := util.NewScannerForInput(2, false)
	defer f.Close()
	safe := 0

	for f.Scan() {
		raw := f.Text()

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
