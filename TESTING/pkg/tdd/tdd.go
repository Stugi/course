package tdd

import "math"

type Point struct {
	X, Y float64
}

func Distance(a, b Point) float64 {
	if a.X < 0 || a.Y < 0 || b.X < 0 || b.Y < 0 {
		return 0
	}
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
}
