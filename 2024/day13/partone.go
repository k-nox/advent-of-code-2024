package day13

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

type machine struct {
	a     image.Point
	b     image.Point
	prize image.Point
}

func (m machine) String() string {
	return fmt.Sprintf("button a %s\nbutton b %s\nprize at %s\n", m.a, m.b, m.prize)
}

func PartOne(useSample bool) int {
	f := helper.OpenInput(2024, 13, useSample)
	defer f.Close()
	machines := parseInp(f)
	tokens := 0
	for _, mach := range machines {
		toWin, possible := mach.play()
		if possible {
			tokens += toWin
		}
	}

	return tokens
}

func parseInp(f *os.File) []machine {
	machines := []machine{}
	buttonRegexp := regexp.MustCompile(`^Button (A|B): X\+(\d+), Y\+(\d+)$`)
	prizeRegexp := regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)$`)

	scanner := bufio.NewScanner(f)

	mach := machine{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		buttonMatches := buttonRegexp.FindAllStringSubmatch(line, -1)
		if buttonMatches != nil {
			buttonLabel := buttonMatches[0][1]
			x, err := strconv.Atoi(buttonMatches[0][2])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(buttonMatches[0][3])
			if err != nil {
				panic(err)
			}

			button := image.Pt(x, y)

			switch buttonLabel {
			case "A":
				mach.a = button
			case "B":
				mach.b = button
			}
			continue
		}

		prizeMatches := prizeRegexp.FindAllStringSubmatch(line, -1)
		if prizeMatches != nil {
			x, err := strconv.Atoi(prizeMatches[0][1])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(prizeMatches[0][2])
			if err != nil {
				panic(err)
			}
			mach.prize = image.Pt(x, y)
			machines = append(machines, mach)
			mach = machine{}
		}
	}
	return machines
}

// returns the # of tokens needed to win; bool is true if game can be won, false if not
func (mach machine) play() (int, bool) {
	ax := float32(mach.a.X)
	ay := float32(mach.a.Y)
	bx := float32(mach.b.X)
	by := float32(mach.b.Y)
	px := float32(mach.prize.X)
	py := float32(mach.prize.Y)

	bTokensFloat := (ay*px - ax*py) / (ay*bx - ax*by)
	aTokensFloat := (px - bx*bTokensFloat) / ax

	aTokens := int(aTokensFloat)
	bTokens := int(bTokensFloat)

	if mach.a.X*aTokens+mach.b.X*bTokens != mach.prize.X {
		return 0, false
	}

	if mach.a.Y*aTokens+mach.b.Y*bTokens != mach.prize.Y {
		return 0, false
	}

	return (aTokens * 3) + (bTokens), true
}
