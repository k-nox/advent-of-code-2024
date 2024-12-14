package day14

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2024, 14, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	max := image.Pt(101, 103)
	if useSample {
		max = image.Pt(11, 7)
	}

	bounds := image.Rect(0, 0, max.X, max.Y)
	// quads := quads(bounds)

	robots := []*robot{}
	uniqueRobots := helper.NewSet[image.Point]()
	for scanner.Scan() {
		bot := parseRobot(strings.TrimSpace(scanner.Text()))
		robots = append(robots, bot)
		uniqueRobots.Add(bot.position)
	}

	for i := 0; i < 10000; i++ {
		uniqueRobots = moveRobots(robots, bounds)
		file := filepath.Join("output", "2024", "day14", fmt.Sprintf("%d.png", i+1))
		render(bounds, uniqueRobots, file)
	}

	// 7569
	return 0
}

func moveRobots(robots []*robot, bounds image.Rectangle) helper.Set[image.Point] {
	uniqueRobots := helper.NewSet[image.Point]()
	for _, bot := range robots {
		bot.move(bounds)
		uniqueRobots.Add(bot.position)
	}
	return uniqueRobots
}

func render(bounds image.Rectangle, robots helper.Set[image.Point], name string) {
	img := image.NewRGBA(bounds)
	palette := helper.RosePineMoon()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := palette.Base()
			if robots.Contains(image.Pt(x, y)) {
				c = palette.Pine()
			}
			img.Set(x, y, c)
		}
	}
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = png.Encode(f, img); err != nil {
		panic(err)
	}
}

func printArea(bounds image.Rectangle, robots helper.Set[image.Point]) {
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			toPrint := "."
			if robots.Contains(image.Pt(x, y)) {
				toPrint = "x"
			}
			fmt.Print(toPrint)
		}
		fmt.Print("\n")
	}
}
