package displacement

import "github.com/Konstantin8105/GoFea/input/point"

// Dim2 - displacement in 2d
// Base unit for coordinates - meter
type Dim2 struct {
	X, Y  float64
	Index point.Index
}
