package day09

import (
	"container/list"
	"fmt"
	"os"
)

type block struct {
	isFile  bool
	id      int
	size    int
	checked bool
}

func PartTwo(useSample bool) int {
	file := "input"
	if useSample {
		file = "sample"
	}

	f, err := os.ReadFile(fmt.Sprintf("input/day09/%s.txt", file))
	if err != nil {
		panic(err)
	}

	inp := string(f)
	disk := parseExpanded(inp)
	moveBlocksSizeAware(disk)
	return checksum(toArray(disk))
}

func moveBlocksSizeAware(disk *list.List) {
	for right := disk.Back(); right != nil; right = right.Prev() {
		rightBlock := right.Value.(block)
		if !rightBlock.isFile || rightBlock.checked {
			continue
		}

		rightBlock.checked = true

		for left := disk.Front(); left != right; left = left.Next() {
			leftBlock := left.Value.(block)
			if leftBlock.isFile || leftBlock.size < rightBlock.size {
				continue
			}

			if leftBlock.size > rightBlock.size {
				disk.InsertAfter(block{
					isFile: false,
					size:   leftBlock.size - rightBlock.size,
				}, left)
			}

			left.Value = block{
				isFile: true,
				size:   rightBlock.size,
				id:     rightBlock.id,
			}
			right.Value = block{
				isFile: false,
				size:   rightBlock.size,
			}
			break
		}
	}
}

func parseExpanded(inp string) *list.List {
	l := list.New()
	isFile := true
	id := 0
	for _, char := range inp {
		l.PushBack(block{
			isFile: isFile,
			id:     id,
			size:   int(char - '0'),
		})
		if isFile {
			id++
		}
		isFile = !isFile
	}
	return l
}

func toArray(disk *list.List) []*int {
	diskArr := []*int{}
	for e := disk.Front(); e != nil; e = e.Next() {
		b := e.Value.(block)
		var id *int
		if b.isFile {
			id = &b.id
		}
		for range b.size {
			diskArr = append(diskArr, id)
		}
	}
	return diskArr
}
