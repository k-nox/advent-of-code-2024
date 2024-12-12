package day12

import (
	"image"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2024, 12, useSample)
	defer f.Close()

	garden := parseInp(f)

	cost := 0
	regions := traverseAllRegions2(garden)
	for _, r := range regions {
		cost += r.price()
	}

	return cost
}

func traverseAllRegions2(garden grid) []region {
	regions := []region{}
	visited := helper.NewSet[image.Point]()
	for plot := range garden {
		if visited.Contains(plot) {
			continue
		}
		var reg region
		reg, visited = traverseRegion2(garden, plot, visited, reg)
		regions = append(regions, reg)
	}
	return regions
}

func traverseRegion2(garden grid, plot image.Point, visited helper.Set[image.Point], reg region) (region, helper.Set[image.Point]) {
	plant := garden[plot]
	reg.area++
	visited.Add(plot)
	reg.perimeter += corners(garden, plot)

	for _, dir := range helper.CardinalDirections {
		next := dir(plot)
		if nextPlant, valid := garden[next]; !valid || nextPlant != plant || visited.Contains(next) {
			continue
		}
		reg, visited = traverseRegion2(garden, next, visited, reg)
	}
	return reg, visited
}

func corners(garden grid, plot image.Point) int {
	corners := 0
	plant := garden[plot]
	upIsEdge := isEdge(garden, plant, helper.Up(plot))
	rightIsEdge := isEdge(garden, plant, helper.Right(plot))
	downIsEdge := isEdge(garden, plant, helper.Down(plot))
	leftIsEdge := isEdge(garden, plant, helper.Left(plot))
	upRightIsEdge := isEdge(garden, plant, helper.UpRight(plot))
	upLeftIsEdge := isEdge(garden, plant, helper.UpLeft(plot))
	downRightIsEdge := isEdge(garden, plant, helper.DownRight(plot))
	downLeftIsEdge := isEdge(garden, plant, helper.DownLeft(plot))

	if upIsEdge && leftIsEdge {
		corners++
	}

	if downIsEdge && leftIsEdge {
		corners++
	}

	if upIsEdge && rightIsEdge {
		corners++
	}

	if downIsEdge && rightIsEdge {
		corners++
	}

	if !upIsEdge && !rightIsEdge && upRightIsEdge {
		corners++
	}

	if !downIsEdge && !leftIsEdge && downLeftIsEdge {
		corners++
	}

	if !upIsEdge && !leftIsEdge && upLeftIsEdge {
		corners++
	}

	if !downIsEdge && !rightIsEdge && downRightIsEdge {
		corners++
	}

	return corners
}

func isEdge(garden grid, plant rune, plot image.Point) bool {
	val, ok := garden[plot]
	return !ok || val != plant
}
