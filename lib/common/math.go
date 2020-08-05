package common

func Lerp(x, y, t float64) float64 {
	return x + (y-x)*t
}
