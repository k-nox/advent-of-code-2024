package parse

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func OpenFile(year int, day int, name string) *os.File {
	p := path(year, day, name)
	file, err := os.Open(p)
	if err != nil {
		panic(err)
	}

	return file
}

func OpenInput(year int, day int, useSample bool) *os.File {
	name := "input.txt"
	if useSample {
		name = "sample.txt"
	}

	return OpenFile(year, day, name)
}

func ReadInput(year int, day int, useSample bool) []byte {
	name := "input.txt"
	if useSample {
		name = "sample.txt"
	}

	return ReadFile(year, day, name)
}

func ReadFile(year int, day int, name string) []byte {
	p := path(year, day, name)

	f, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}

	return f
}

func path(year int, day int, name string) string {
	return filepath.Join("input", strconv.Itoa(year), fmt.Sprintf("day%02d", day), name)
}
