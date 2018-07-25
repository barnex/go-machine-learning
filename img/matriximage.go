package img

import (
	"image"
	"image/color"
	"math"

	vs "github.com/barnex/vectorstream"
)

func FromMatrix(m vs.M, min, max float64) image.Image {
	return &matrixImage{m, min, max}
}

type matrixImage struct {
	m        vs.M
	min, max float64
}

func (m *matrixImage) At(i, j int) color.Color {
	y := srgb(m.m.Elem(j)[i])
	return color.Gray{uint8(y * 255)}
}

func (m *matrixImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.m.Size[0], m.m.Size[1])
}

func (m *matrixImage) ColorModel() color.Model {
	return nil
}

// linear to sRGB conversion
// https://en.wikipedia.org/wiki/SRGB
func srgb(c float64) float64 {
	c = clip(c)
	if c <= 0.0031308 {
		return 12.92 * c
	}
	c = 1.055*math.Pow(float64(c), 1./2.4) - 0.05
	if c > 1 {
		return 1
	}
	return c
}

// clip color value between 0 and 1
func clip(v float64) float64 {
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	return v
}
