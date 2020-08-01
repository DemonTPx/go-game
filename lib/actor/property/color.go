package property

import "fmt"

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

func NewColor(r float64, g float64, b float64, a float64) Color {
	return Color{r, g, b, a}
}

func (c *Color) String() string {
	return fmt.Sprintf("<Color red=%f green=%f blue=%f alpha=%f>", c.R, c.G, c.B, c.A)
}
