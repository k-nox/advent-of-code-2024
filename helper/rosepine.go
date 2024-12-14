package helper

import "image/color"

type Palette interface {
	Base() color.Color
	Surface() color.Color
	Overlay() color.Color
	Muted() color.Color
	Subtle() color.Color
	Text() color.Color
	Love() color.Color
	Gold() color.Color
	Rose() color.Color
	Pine() color.Color
	Foam() color.Color
	Iris() color.Color
	HighlightLow() color.Color
	HighlightMed() color.Color
	HighlightHigh() color.Color
}

type rosePineMoon struct{}

// var RosePineMoon Palette = rosePineMoon{}
func RosePineMoon() Palette {
	return rosePineMoon{}
}

func (rosePineMoon) Base() color.Color {
	return color.RGBA{35, 33, 54, 0xff}
}

func (rosePineMoon) Surface() color.Color {
	return color.RGBA{42, 39, 63, 0xff}
}

func (rosePineMoon) Overlay() color.Color {
	return color.RGBA{57, 53, 82, 0xff}
}

func (rosePineMoon) Muted() color.Color {
	return color.RGBA{110, 106, 134, 0xff}
}

func (rosePineMoon) Subtle() color.Color {
	return color.RGBA{144, 140, 170, 0xff}
}

func (rosePineMoon) Text() color.Color {
	return color.RGBA{224, 222, 244, 0xff}
}

func (rosePineMoon) Love() color.Color {
	return color.RGBA{235, 111, 146, 0xff}
}

func (rosePineMoon) Gold() color.Color {
	return color.RGBA{246, 193, 119, 0xff}
}

func (rosePineMoon) Rose() color.Color {
	return color.RGBA{234, 154, 151, 0xff}
}

func (rosePineMoon) Pine() color.Color {
	return color.RGBA{62, 143, 176, 0xff}
}

func (rosePineMoon) Foam() color.Color {
	return color.RGBA{156, 207, 216, 0xff}
}

func (rosePineMoon) Iris() color.Color {
	return color.RGBA{196, 167, 231, 0xff}
}

func (rosePineMoon) HighlightLow() color.Color {
	return color.RGBA{42, 40, 62, 0xff}
}

func (rosePineMoon) HighlightMed() color.Color {
	return color.RGBA{68, 65, 90, 0xff}
}

func (rosePineMoon) HighlightHigh() color.Color {
	return color.RGBA{86, 82, 110, 0xff}
}
