package calculator

import "math"

func CalculateDistance(a, b, c, d float64) float64 {
	x := math.Pow(a-c, 2)
	y := math.Pow(b-d, 2)
	return math.Sqrt(x + y)
}
