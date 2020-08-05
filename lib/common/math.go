package common

func Lerp(x, y, t float64) float64 {
	return x + (y-x)*t
}

func Clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
