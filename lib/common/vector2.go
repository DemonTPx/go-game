package common

import "fmt"

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x float64, y float64) Vector2 {
	return Vector2{x, y}
}

func (v *Vector2) String() string {
	return fmt.Sprintf("<Vector2 x=%f y=%f>", v.X, v.Y)
}

func (v *Vector2) Add(o *Vector2) Vector2 {
	return Vector2{X: v.X + o.X, Y: v.Y + o.Y}
}

func (v *Vector2) Sub(o *Vector2) Vector2 {
	return Vector2{X: v.X - o.X, Y: v.Y - o.Y}
}

func (v *Vector2) Multi(o *Vector2) Vector2 {
	return Vector2{X: v.X * o.X, Y: v.Y * o.Y}
}

func (v *Vector2) MultiFloat64(m float64) Vector2 {
	return Vector2{X: v.X * m, Y: v.Y * m}
}
