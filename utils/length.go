package utils

import (
	"math"

	"github.com/Konstantin8105/GoFea/input/point"
)

// LengthDim2 - distance between 2 point in 2d
func LengthDim2(p0, p1 point.Dim2) float64 {
	return math.Sqrt(math.Pow(p0.X-p1.X, 2.0) + math.Pow(p0.Y-p1.Y, 2.0))
}
