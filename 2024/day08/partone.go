package day08

import (
	"bufio"
	"image"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

type grid map[image.Point]string

func PartOne(useSample bool) int {
	f := helper.OpenInput(2024, 8, useSample)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	antinodes := map[image.Point]bool{}
	grid, antennas := parseInp(scanner)
	for _, locations := range antennas {
		antinodesForLoc := findAntinodes(grid, locations)
		for _, n := range antinodesForLoc {
			if _, ok := antinodes[n]; !ok {
				antinodes[n] = true
			}
		}
	}

	return len(antinodes)
}

func parseInp(scanner *bufio.Scanner) (grid, map[string][]image.Point) {
	g := grid{}
	antennas := map[string][]image.Point{}

	y := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		for x := range line {
			pt := image.Point{
				X: x,
				Y: y,
			}
			char := string(line[x])
			g[pt] = char
			if char != "." {
				antennas[char] = append(antennas[char], pt)
			}
		}
		y++
	}
	return g, antennas
}

func findAntinodesForPair(a image.Point, b image.Point) (image.Point, image.Point) {
	distance := a.Sub(b)
	return a.Add(distance), b.Sub(distance)
}

func findAntinodes(g grid, locations []image.Point) []image.Point {
	antinodes := []image.Point{}
	for i := 0; i < len(locations)-1; i++ {
		for j := 1; j < len(locations); j++ {
			if i == j {
				continue
			}
			antinodeA, antinodeB := findAntinodesForPair(locations[i], locations[j])
			if _, ok := g[antinodeA]; ok {
				antinodes = append(antinodes, antinodeA)
			}
			if _, ok := g[antinodeB]; ok {
				antinodes = append(antinodes, antinodeB)
			}
		}
	}
	return antinodes
}
