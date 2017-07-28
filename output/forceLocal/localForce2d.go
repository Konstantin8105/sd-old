package forceLocal

import "github.com/Konstantin8105/GoFea/input/element"

// Dim2 - local forces in point
// Unit - N
type Dim2 struct {
	Fx, Fy, M float64
}

// Beam2d - local forces for beam 2d
// Unit - N
type BeamDim2 struct {
	Begin, End Dim2
	Index      element.Index
}
