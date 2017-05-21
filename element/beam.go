package element

import "github.com/Konstantin8105/GoFea/point"

// Beam - property of element
type Beam struct {
	Index        BeamIndex
	PointIndexes [2]point.Index
}
