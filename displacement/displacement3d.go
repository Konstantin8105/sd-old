package displacement

import "github.com/Konstantin8105/GoFea/point"

// Dim3 - displacement in 3d
// Base unit for coordinates - meter
type Dim3 struct {
	X, Y, Z float64
	Index   point.Index
}
