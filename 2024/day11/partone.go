package day11

import (
	"bufio"
	"os"
	"strconv"

	"github.com/k-nox/advent-of-code-solutions/parse"
)

func PartOne(useSample bool) int {
	f := parse.OpenInput(2024, 11, useSample)
	defer f.Close()

	stones := parseInp(f)
	for range 25 {
		stones = blink(stones)
	}

	return len(stones)
}

func parseInp(f *os.File) []int {
	stones := []int{}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		stone, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		stones = append(stones, stone)
	}

	return stones
}

func blink(stones []int) []int {
	currentLength := len(stones)
	for i := 0; i < currentLength; i++ {
		if stones[i] == 0 {
			stones[i] = 1
			continue
		}

		stone := strconv.Itoa(stones[i])
		if len(stone)%2 == 0 {
			firstHalf, err := strconv.Atoi(stone[:len(stone)/2])
			if err != nil {
				panic(err)
			}
			secondHalf, err := strconv.Atoi(stone[len(stone)/2:])
			if err != nil {
				panic(err)
			}
			stones[i] = firstHalf
			stones = append(stones, secondHalf)
			continue
		}

		stones[i] *= 2024
	}
	return stones
}
