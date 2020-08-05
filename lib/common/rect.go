package common

type Rect struct {
	X, Y, W, H float64
}

func NewRect(x, y, w, h float64) Rect {
	return Rect{X: x, Y: y, W: w, H: h}
}

func (r *Rect) X2() float64 {
	return r.X + r.W
}

func (r *Rect) Y2() float64 {
	return r.Y + r.H
}
