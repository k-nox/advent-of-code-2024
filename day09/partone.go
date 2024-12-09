package day09

import (
	"fmt"
	"os"
)

func PartOne(useSample bool) int {
	file := "input"
	if useSample {
		file = "sample"
	}

	f, err := os.ReadFile(fmt.Sprintf("input/day09/%s.txt", file))
	if err != nil {
		panic(err)
	}

	inp := string(f)
	disk := parse(inp)
	moveBlocks(disk)

	return checksum(disk)
}

func parse(inp string) []*int {
	disk := []*int{}
	isFile := true
	id := 0
	for _, char := range inp {
		fileId := id
		val := int(char - '0')
		var b *int
		if isFile {
			b = &fileId
		}
		for range val {
			disk = append(disk, b)
		}
		if isFile {
			id++
		}
		isFile = !isFile
	}
	return disk
}

func moveBlocks(disk []*int) {
	left := 0
	right := len(disk) - 1
	for left != right {
		if disk[left] != nil {
			left++
			continue
		}

		if disk[right] == nil {
			right--
			continue
		}

		disk[left] = disk[right]
		disk[right] = nil
		left++
		right--
	}
}

func checksum(disk []*int) int {
	sum := 0
	for idx, id := range disk {
		if id == nil {
			continue
		}
		sum += (idx * *id)
	}
	return sum
}

func printDisk(disk []*int) {
	fmt.Printf("[ ")
	for _, d := range disk {
		if d == nil {
			fmt.Printf("%v ", d)
		} else {
			fmt.Printf("%d ", *d)
		}
	}
	fmt.Printf(" ]\n")
}
