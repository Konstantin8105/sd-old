package displacement

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/point"
)

type Dim2 struct {
	Dx, Dy, Dm float64
}

// Dim2 - displacement in 2d
// Base unit for coordinates - meter
type PointDim2 struct {
	Dim2
	Index point.Index
}

type BeamDim2 struct {
	Begin, End Dim2
	Index      element.Index
}
