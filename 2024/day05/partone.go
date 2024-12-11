package day05

import (
	"strconv"
	"strings"

	"github.com/k-nox/aoc/util"
)

func PartOne(useSample bool) int {
	f := util.NewScannerForInput(2024, 5, useSample)
	defer f.Close()

	// X|Y
	rules := map[string][]string{} // key = X, val = Y; x must come before y
	parsingRules := true
	sep := "|"
	sum := 0
	for f.Scan() {
		curr := strings.TrimSpace(f.Text())

		if curr == "" {
			parsingRules = false
			sep = ","
			continue
		}

		raw := strings.Split(curr, sep)

		if parsingRules {
			x := raw[0]
			y := raw[1]
			rules[x] = append(rules[x], y)
			continue
		}

		seen := map[string]bool{}
		ruleBroken := false
		for _, page := range raw {
			seen[page] = true
			if rule, ok := rules[page]; ok {
				for _, y := range rule {
					if _, ok := seen[y]; ok {
						ruleBroken = true
						break
					}
				}
				if ruleBroken {
					break
				}
			}

		}
		if !ruleBroken {
			mid, err := strconv.Atoi(raw[len(raw)/2])
			if err != nil {
				panic(err)
			}
			sum += mid
		}
	}
	return sum
}
