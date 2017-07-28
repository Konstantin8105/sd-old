package point

import "math"

// Dim2 - position of point in 2d (dimentions)
// Base unit for coordinates - meter
type Dim2 struct {
	X, Y  float64
	Index Index
}

// LengthDim2 - distance between 2 points in 2d
func LengthDim2(points [2]Dim2) float64 {
	return math.Sqrt(math.Pow(points[0].X-points[1].X, 2.0) + math.Pow(points[0].Y-points[1].Y, 2.0))
}
