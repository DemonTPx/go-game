package common

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Color struct {
	R, G, B, A float64
}

func NewColor(r, g, b, a float64) Color {
	return Color{R: r, G: g, B: b, A: a}
}

func NewColorWhite() Color {
	return NewColor(1, 1, 1, 1)
}

func (c *Color) String() string {
	return fmt.Sprintf("<Color red=%f green=%f blue=%f alpha=%f>", c.R, c.G, c.B, c.A)
}

func (c *Color) Lerp(o *Color, t float64) Color {
	return Color{
		R: Lerp(c.R, o.R, t),
		G: Lerp(c.G, o.G, t),
		B: Lerp(c.B, o.B, t),
		A: Lerp(c.A, o.A, t),
	}
}

func (c *Color) ToSDL() sdl.Color {
	return sdl.Color{
		R: uint8(c.R * 255),
		G: uint8(c.G * 255),
		B: uint8(c.B * 255),
		A: uint8(c.A * 255),
	}
}
