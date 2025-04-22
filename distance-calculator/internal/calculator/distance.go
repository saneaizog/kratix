package calculator

import "math"

// CalculateDistance вычисляет расстояние между двумя точками на плоскости
// Принимает координаты двух точек (x1, y1) и (x2, y2)
func CalculateDistance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}
