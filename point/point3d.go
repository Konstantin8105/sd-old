package point

import "math"

// Dim3 - position of point in 3d (dimentions)
// Base unit for coordinates - meter
type Dim3 struct {
	X, Y, Z float64
	Index   Index
}

// LenghtDim3 - distance between 2 points in 3d
func LenghtDim3(points [2]Dim3) float64 {
	return math.Sqrt(math.Pow(points[0].X-points[1].X, 2.0) + math.Pow(points[0].Y-points[1].Y, 2.0) + math.Pow(points[0].Z-points[1].Z, 2.0))
}
