package day05

import (
	"bufio"
	"regexp"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
	f := helper.OpenInput(2015, 5, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	nice := 0
	for scanner.Scan() {
		if check(scanner.Text()) {
			nice++
		}
	}
	return nice
}

func check(s string) bool {
	mustNotRe := regexp.MustCompile(`ab|cd|pq|xy`)
	if mustNotRe.MatchString(s) {
		return false
	}

	vowels := 0
	foundDouble := false
	for idx, char := range s {
		if !foundDouble && idx > 0 && char == rune(s[idx-1]) {
			foundDouble = true
		}
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
	}
	return vowels >= 3 && foundDouble
}
