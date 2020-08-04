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

func (c *Color) String() string {
	return fmt.Sprintf("<Color red=%f green=%f blue=%f alpha=%f>", c.R, c.G, c.B, c.A)
}

func (c *Color) ToSDL() sdl.Color {
	return sdl.Color{
		R: uint8(c.R * 255),
		G: uint8(c.G * 255),
		B: uint8(c.B * 255),
		A: uint8(c.A * 255),
	}
}
