package day14

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

type robot struct {
	position image.Point
	velocity image.Point
}

func (r robot) String() string {
	return fmt.Sprintf("robot position = %s, velocity = %s", r.position, r.velocity)
}

type quad struct {
	bounds image.Rectangle
	count  int
}

func PartOne(useSample bool) int {
	f := helper.OpenInput(2024, 14, useSample)
	defer f.Close()

	max := image.Pt(101, 103)
	if useSample {
		max = image.Pt(11, 7)
	}

	bounds := image.Rect(0, 0, max.X, max.Y)
	quads := quads(bounds)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		robot := parseRobot(strings.TrimSpace(scanner.Text()))
		for range 100 {
			robot.move(bounds)
		}
		for idx := range quads {
			if robot.position.In(quads[idx].bounds) {
				quads[idx].count++
			}
		}
	}

	safetyFactor := 1
	for _, quad := range quads {
		safetyFactor *= quad.count
	}

	return safetyFactor
}

func quads(bounds image.Rectangle) []quad {
	mid := bounds.Max.Div(2)

	return []quad{
		{
			bounds: image.Rect(0, 0, mid.X, mid.Y),
		},
		{
			bounds: image.Rect(mid.X+1, 0, bounds.Max.X, mid.Y),
		},
		{
			bounds: image.Rect(0, mid.Y+1, mid.X, bounds.Max.Y),
		},
		{
			bounds: image.Rect(mid.X+1, mid.Y+1, bounds.Max.X, bounds.Max.Y),
		},
	}
}

func parseRobot(line string) *robot {
	r := robot{}
	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	match := re.FindStringSubmatch(line)
	r.position = image.Pt(strToNum(match[1]), strToNum(match[2]))
	r.velocity = image.Pt(strToNum(match[3]), strToNum(match[4]))

	return &r
}

func strToNum(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func (r *robot) move(bounds image.Rectangle) {
	pos := r.position.Add(r.velocity)
	if pos.In(bounds) {
		r.position = pos
		return
	}

	// need to wrap around...
	r.moveX(bounds)
	r.moveY(bounds)
}

func (r *robot) moveX(bounds image.Rectangle) {
	velocity := image.Pt(1, 0)
	wrap := bounds.Min.X
	if r.velocity.X < 0 {
		velocity = image.Pt(-1, 0)
		wrap = bounds.Max.X - 1
	}

	for range int(math.Abs(float64(r.velocity.X))) {
		next := r.position.Add(velocity)
		if !next.In(bounds) {
			next = image.Pt(wrap, r.position.Y)
		}
		r.position = next
	}
}

func (r *robot) moveY(bounds image.Rectangle) {
	velocity := image.Pt(0, 1)
	wrap := bounds.Min.Y
	if r.velocity.Y < 0 {
		velocity = image.Pt(0, -1)
		wrap = bounds.Max.Y - 1
	}

	for range int(math.Abs(float64(r.velocity.Y))) {
		next := r.position.Add(velocity)
		if !next.In(bounds) {
			next = image.Pt(r.position.X, wrap)
		}
		r.position = next
	}
}
