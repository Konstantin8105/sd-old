package element

import "github.com/Konstantin8105/GoFea/point"

// BeamIndex - alias
type BeamIndex int

// Beam - property of element
type Beam struct {
	Index       BeamIndex
	PointIndexs [2]point.Index
}
