package day05

import (
	"bufio"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	file := "input.txt"
	if useSample {
		file = "sample2.txt"
	}
	f := helper.OpenFile(2015, 5, file)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	nice := 0
	for scanner.Scan() {
		if check2(scanner.Text()) {
			nice++
		}
	}

	return nice
}

// It contains a pair of any two letters that appears at least twice
// in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa),
// but not like aaa (aa, but it overlaps).

// It contains at least one letter which repeats with exactly one letter
// between them, like xyx, abcdefeghi (efe), or even aaa.

func check2(s string) bool {
	return checkForRepeatedPair(s) && checkRepeatedLetterWithSpace(s)
}

func checkForRepeatedPair(s string) bool {
	for i := 1; i < len(s); i++ {
		pair := string(s[i-1]) + string(s[i])
		count := strings.Count(s, pair)
		if count > 1 {
			return true
		}
	}
	return false
}

func checkRepeatedLetterWithSpace(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			return true
		}
	}
	return false
}
