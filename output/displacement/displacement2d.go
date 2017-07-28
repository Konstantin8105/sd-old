package displacement

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/point"
)

// Dim2 - displacement in 2d
// Base unit for coordinates - meter
type Dim2 struct {
	Dx, Dy, Dm float64
}

// PointDim2 - displacement of point in 2d
// Base unit for coordinates - meter
type PointDim2 struct {
	Dim2
	Index point.Index
}

// BeamDim2 - displacement of beam points in 2d
// Base unit for coordinates - meter
type BeamDim2 struct {
	Begin, End Dim2
	Index      element.Index
}
