package forceLocal

import "github.com/Konstantin8105/GoFea/element"

// Beam2d - local forces for beam 2d
// Unit - N
type Beam2d struct {
	Fx, Fy float64
	M      float64
	Index  element.BeamIndex
}
