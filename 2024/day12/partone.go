package day12

import (
	"bufio"
	"fmt"
	"image"
	"os"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

type grid map[image.Point]rune

type region struct {
	area      int
	perimeter int
}

func (r region) price() int {
	return r.area * r.perimeter
}

func (r region) String() string {
	return fmt.Sprintf("region.area = %d\nregion.perimeter = %d\n", r.area, r.perimeter)
}

func PartOne(useSample bool) int {
	f := helper.OpenInput(2024, 12, useSample)
	defer f.Close()

	garden := parseInp(f)
	regions := traverseAllRegions(garden)
	cost := 0
	for _, r := range regions {
		cost += r.price()
	}

	return cost
}

func parseInp(f *os.File) grid {
	scanner := bufio.NewScanner(f)
	garden := grid{}

	y := 0
	for scanner.Scan() {
		row := scanner.Text()
		for x, char := range row {
			garden[image.Pt(x, y)] = char
		}
		y++
	}
	return garden
}

func traverseAllRegions(garden grid) []region {
	regions := []region{}
	visited := helper.NewSet[image.Point]()
	for plot := range garden {
		if visited.Contains(plot) {
			continue
		}
		var reg region
		reg, visited = traverseRegion(garden, plot, visited, reg)
		regions = append(regions, reg)
	}
	return regions
}

func traverseRegion(garden grid, plot image.Point, visited helper.Set[image.Point], reg region) (region, helper.Set[image.Point]) {
	plant := garden[plot]
	reg.area++
	visited.Add(plot)
	for _, dir := range helper.CompassDirections {
		next := dir(plot)
		if _, valid := garden[next]; !valid {
			reg.perimeter++
			continue
		}

		if garden[next] != plant {
			reg.perimeter++
			continue
		}

		if visited.Contains(next) {
			continue
		}

		reg, visited = traverseRegion(garden, next, visited, reg)
	}

	return reg, visited
}
