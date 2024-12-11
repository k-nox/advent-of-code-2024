package day11

import (
	"bufio"
	"maps"
	"os"
	"strconv"

	"github.com/k-nox/advent-of-code-solutions/parse"
)

func PartTwo(useSample bool) int {
	f := parse.OpenInput(2024, 11, useSample)
	defer f.Close()

	stones := parseInp2(f)

	for range 75 {
		stones = blink2(stones)
	}
	sum := 0

	for _, count := range stones {
		sum += count
	}

	return sum
}

func parseInp2(f *os.File) map[int]int {
	stones := map[int]int{}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		stone, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		stones[stone]++
	}
	return stones
}

func blink2(stones map[int]int) map[int]int {
	stonecopy := maps.Clone(stones)
	for stone, count := range stones {
		if count == 0 {
			continue
		}
		if stone == 0 {
			stonecopy[1] += count
			stonecopy[0] -= count
			continue
		}
		stoneStr := strconv.Itoa(stone)
		if len(stoneStr)%2 == 0 {
			middle := len(stoneStr) / 2
			firstHalf, _ := strconv.Atoi(stoneStr[:middle])
			secondHalf, _ := strconv.Atoi(stoneStr[middle:])
			stonecopy[firstHalf] += count
			stonecopy[secondHalf] += count
			stonecopy[stone] -= count
			continue
		}

		stonecopy[stone*2024] += count
		stonecopy[stone] -= count

	}
	return stonecopy
}

// func blinkSingleStone(stone int, count int, memo map[int][]int) (int, map[int][]int) {
// 	sum := 0
// 	for i := count; i > 0; i-- {
// 		if becomes, ok := memo[stone]; ok {
// 			stone = becomes[0]
// 			if len(becomes) > 1 {
// 				var n int
// 				n, memo = blinkSingleStone(becomes[1], i-1, memo)
// 				sum += n + 1
// 			}
// 			continue
// 		}

// 		stoneStr := strconv.Itoa(stone)
// 		if len(stoneStr)%2 == 0 {
// 			firstHalf, err := strconv.Atoi(stoneStr[:len(stoneStr)/2])
// 			if err != nil {
// 				panic(err)
// 			}
// 			secondHalf, err := strconv.Atoi(stoneStr[len(stoneStr)/2:])
// 			if err != nil {
// 				panic(err)
// 			}
// 			memo[stone] = []int{firstHalf, secondHalf}
// 			stone = firstHalf
// 			var n int
// 			n, memo = blinkSingleStone(secondHalf, i-1, memo)
// 			sum += n + 1
// 			continue
// 		}

// 		newStone := stone * 2024
// 		memo[stone] = []int{newStone}
// 		stone = newStone
// 	}

// 	return sum, memo
// }
