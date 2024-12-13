package day06

import (
	"bufio"
	"image"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2015, 6, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	instructions := []instruction{}
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}

	lights := map[image.Point]int{}
	for _, instr := range instructions {
		lights = instr.perform2(lights)
	}

	total := 0
	for _, brightness := range lights {
		total += brightness
	}

	return total
}

func (instr instruction) perform2(lights map[image.Point]int) map[image.Point]int {
	for x := instr.start.X; x <= instr.end.X; x++ {
		for y := instr.start.Y; y <= instr.end.Y; y++ {
			light := image.Pt(x, y)
			switch instr.action {
			case on:
				lights[light]++
			case off:
				if lights[light] > 0 {
					lights[light]--
				}
			case toggle:
				lights[light] += 2
			}
		}
	}
	return lights
}
