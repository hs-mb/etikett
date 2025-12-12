package eplutil

import (
	"fmt"
	"image"
	"image/color"
)

func imageToBytes(img image.Image) ([]byte, int, int) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	byteWidth := width / 8
	if width % 8 != 0 {
		byteWidth += 1
	}
	
	data := make([]byte, byteWidth * height)

	for y := range height {
		for byteN := range byteWidth {
			for offset := range 8 {
				x := byteN * 8 + offset
				byt := byte(0b1000_0000) >> offset
				luma := uint8(255)
				if x < width {
					r, g, b, _ := img.At(x, y).RGBA()
					luma, _, _ = color.RGBToYCbCr(uint8(r >> 8), uint8(g >> 8), uint8(b >> 8))
				}
				if luma > 128 {
					data[y * byteWidth + byteN] |= byt
				} else {
					data[y * byteWidth + byteN] &= ^byt
				}
			}
		}
	}

	return data, byteWidth, height
}

func (b *EPLBuilder) Image(x, y int, img image.Image) {
	data, lineBytes, lines := imageToBytes(img)
	b.ImageBytes(x, y, lineBytes, lines, data)
}

func (b *EPLBuilder) ImageBytes(x, y, lineBytes, lines int, data []byte) {
	b.WriteString(fmt.Sprintf("GW%d,%d,%d,%d,", x, y, lineBytes, lines) + string(data))
}
