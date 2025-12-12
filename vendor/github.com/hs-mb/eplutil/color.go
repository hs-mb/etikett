package eplutil

import "image/color"

type Color uint
const (
	NONE Color = iota
	BLACK
	WHITE
)

func (c Color) Color() color.Color {
	switch c {
	case BLACK:
		return color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		}
	case WHITE:
		return color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}
	}
	panic("Color unset")
}

