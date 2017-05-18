package element

import (
	coordinate "github.com/Konstantin8105/GoFea/Coordinate"
)

// BeamIndex - alias
type BeamIndex int

// Beam - property of element
type Beam struct {
	Index       BeamIndex
	PointIndexs [2]coordinate.PointIndex
}
