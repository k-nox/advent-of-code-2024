package day05

import (
	"slices"
	"strconv"
	"strings"

	"github.com/k-nox/aoc/util"
)

func PartTwo(useSample bool) int {
	f := util.NewScannerForInput(5, useSample)
	defer f.Close()

	rules := map[string][]string{}
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

		if isCorrectUpdate(rules, raw) {
			continue
		}

		slices.SortFunc(raw, func(a, b string) int {
			if rule, ok := rules[a]; ok && slices.Contains(rule, b) {
				return -1
			}
			if rule, ok := rules[b]; ok && slices.Contains(rule, a) {
				return 1
			}
			return 0
		})

		mid, err := strconv.Atoi(raw[len(raw)/2])
		if err != nil {
			panic(err)
		}
		sum += mid
	}

	return sum
}

func isCorrectUpdate(rules map[string][]string, update []string) bool {
	seen := map[string]bool{}
	for _, page := range update {
		seen[page] = true
		if rule, ok := rules[page]; ok {
			for _, y := range rule {
				if _, ok := seen[y]; ok {
					return false
				}
			}
		}
	}
	return true
}
