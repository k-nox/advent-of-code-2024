package helper

import (
	"image"
)

var CardinalDirections = []func(image.Point) image.Point{
	Up,
	Right,
	Down,
	Left,
}

func Left(p image.Point) image.Point {
	return p.Add(image.Pt(-1, 0))
}

func Right(p image.Point) image.Point {
	return p.Add(image.Pt(1, 0))
}

func Up(p image.Point) image.Point {
	return p.Add(image.Pt(0, -1))
}

func Down(p image.Point) image.Point {
	return p.Add(image.Pt(0, 1))
}

func UpLeft(p image.Point) image.Point {
	return Up(Left(p))
}

func UpRight(p image.Point) image.Point {
	return Up(Right(p))
}

func DownLeft(p image.Point) image.Point {
	return Down(Left(p))
}

func DownRight(p image.Point) image.Point {
	return Down(Right(p))
}
