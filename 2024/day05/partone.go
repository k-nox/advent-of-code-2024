package day05

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/parse"
)

func PartOne(useSample bool) int {
	f := parse.OpenInput(2024, 5, useSample)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// X|Y
	rules := map[string][]string{} // key = X, val = Y; x must come before y
	parsingRules := true
	sep := "|"
	sum := 0
	for scanner.Scan() {
		curr := strings.TrimSpace(scanner.Text())

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
