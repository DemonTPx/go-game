package property

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

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func NewVector3(x float64, y float64, z float64) Vector3 {
	return Vector3{x, y, z}
}

func (v *Vector3) String() string {
	return fmt.Sprintf("<Vector3 x=%f y=%f z=%f>", v.X, v.Y, v.Z)
}

func (v *Vector3) Add(o *Vector3) Vector3 {
	return Vector3{X: v.X + o.X, Y: v.Y + o.Y, Z: v.Z + o.Z}
}

func (v *Vector3) Sub(o *Vector3) Vector3 {
	return Vector3{X: v.X - o.X, Y: v.Y - o.Y, Z: v.Z - o.Z}
}

func (v *Vector3) Multi(o *Vector3) Vector3 {
	return Vector3{X: v.X * o.X, Y: v.Y * o.Y, Z: v.Z * o.Z}
}
