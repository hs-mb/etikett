package eplutil

import (
	"image"
	"image/draw"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	_ "embed"
)

//go:embed res/liberation_serif.ttf
var defaultFontData []byte
var defaultFont *opentype.Font

func loadFace(textFont *opentype.Font, size float64) (font.Face, error) {
	face, err := opentype.NewFace(textFont, &opentype.FaceOptions{Size: size, DPI: 72})
	if err != nil {
		return nil, err
	}
	return face, nil
}

type FittedTextOptions struct {
	Font *opentype.Font
	LineSpace int
	CenterX bool
	CenterY bool
	TextColor Color
	BGColor Color
}

func (o FittedTextOptions) font() (*opentype.Font, error) {
	if o.Font == nil {
		if defaultFont == nil {
			var err error
			defaultFont, err = opentype.Parse(defaultFontData)
			if err != nil {
				return nil, err
			}
		}
		return defaultFont, nil
	}
	return o.Font, nil
}

func (o FittedTextOptions) lineSpace() int {
	return o.LineSpace
}

func (o FittedTextOptions) centerX() bool {
	return o.CenterX
}

func (o FittedTextOptions) centerY() bool {
	return o.CenterY
}

func (o FittedTextOptions) textColor() Color {
	if o.TextColor == NONE {
		return BLACK
	}
	return o.TextColor
}

func (o FittedTextOptions) bgColor() Color {
	if o.BGColor == NONE {
		return WHITE
	}
	return o.BGColor
}

func fittedTextImage(text string, width, height int, opts FittedTextOptions) (_ image.Image, err error) {
	faceHeight := func(face font.Face) int {
		return face.Metrics().CapHeight.Ceil()
	}

	lines := strings.Split(text, "\n")

	img := image.NewGray(image.Rect(0, 0, width, height))
	draw.Draw(img, image.Rect(0, 0, width, height), image.NewUniform(opts.bgColor().Color()), image.Pt(0, 0), draw.Src)

	textFont, err := opts.font()
	if err != nil {
		return
	}

	var testFontSize float64 = 96
	face, err := loadFace(textFont, testFontSize)

	var measuredWidth int
	var maxLineIdx int
	for idx, line := range lines {
		currentWidth := font.MeasureString(face, line).Ceil()
		if idx == 0 || measuredWidth < currentWidth {
			measuredWidth = currentWidth
			maxLineIdx = idx
		}
	}
	measuredHeight := len(lines) * faceHeight(face) + (len(lines) - 1) * opts.lineSpace()

	widthSizeFactor := float64(width) / float64(measuredWidth)
	heightSizeFactor := float64(height) / float64(measuredHeight)
	centerWidth := false
	sizeFactor := widthSizeFactor
	if sizeFactor > heightSizeFactor {
		sizeFactor = heightSizeFactor
		centerWidth = true
	}

	fontSize := testFontSize * sizeFactor
	face, err = loadFace(textFont, fontSize)
	if err != nil {
		return
	}

	realWidth := font.MeasureString(face, lines[maxLineIdx]).Ceil()
	realHeight := len(lines) * faceHeight(face) + (len(lines) - 1) * opts.lineSpace()

	startX := 0
	startY := 0
	if opts.centerX() && centerWidth {
		startX = (width - realWidth) / 2
	}
	if opts.centerY() && !centerWidth {
		startY = (height - realHeight) / 2
	}

	for idx, line := range lines {
		d := &font.Drawer{
			Dst: img,
			Src: image.NewUniform(opts.textColor().Color()),
			Face: face,
			Dot: fixed.P(startX, startY + faceHeight(face) * (idx + 1) + opts.lineSpace() * idx),
		}
		d.DrawString(line)
	}

	return img, nil
}

func (b *EPLBuilder) FittedText(text string, x, y, width, height int, opts FittedTextOptions) error {
	img, err := fittedTextImage(text, width, height, opts)
	if err != nil {
		return err
	}
	b.Image(x, y, img)
	return nil
}
