package actor

import "fmt"

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewColor(r uint8, g uint8, b uint8, a uint8) Color {
	return Color{r, g, b, a}
}

func (c *Color) String() string {
	return fmt.Sprintf("<Color red=%d green=%d blue=%d alpha=%d>", c.R, c.G, c.B, c.A)
}
