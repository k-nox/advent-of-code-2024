package day06

import (
	"bufio"
	"image"
	"regexp"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

const (
	on     = "on"
	off    = "off"
	toggle = "toggle"
)

type instruction struct {
	start  image.Point
	end    image.Point
	action string
}

func PartOne(useSample bool) int {
	f := helper.OpenInput(2015, 6, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	instructions := []instruction{}
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}

	lights := helper.NewSet[image.Point]()
	for _, instr := range instructions {
		lights = instr.perform(lights)
	}

	return len(lights)
}

func parseInstruction(line string) instruction {
	instr := instruction{}
	if strings.HasPrefix(line, "turn on") {
		instr.action = on
	} else if strings.HasPrefix(line, "turn off") {
		instr.action = off
	} else {
		instr.action = toggle
	}

	re := regexp.MustCompile(`(\d+),(\d+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	start := matches[0]
	end := matches[1]

	startX, err := strconv.Atoi(start[1])
	if err != nil {
		panic(err)
	}
	startY, err := strconv.Atoi(start[2])
	if err != nil {
		panic(err)
	}

	endX, err := strconv.Atoi(end[1])
	if err != nil {
		panic(err)
	}
	endY, err := strconv.Atoi(end[2])
	if err != nil {
		panic(err)
	}

	instr.start = image.Pt(startX, startY)
	instr.end = image.Pt(endX, endY)

	return instr
}

func (instr instruction) perform(lights helper.Set[image.Point]) helper.Set[image.Point] {
	for x := instr.start.X; x <= instr.end.X; x++ {
		for y := instr.start.Y; y <= instr.end.Y; y++ {
			light := image.Pt(x, y)
			switch instr.action {
			case on:
				lights.Add(light)
			case off:
				lights.Remove(light)
			case toggle:
				if lights.Contains(light) {
					lights.Remove(light)
				} else {
					lights.Add(light)
				}
			}
		}
	}
	return lights
}
