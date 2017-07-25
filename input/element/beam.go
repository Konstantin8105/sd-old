package element

import (
	"github.com/Konstantin8105/GoFea/input/point"
)

// Beam - property of beam element
type Beam struct {
	Index        ElementIndex
	PointIndexes [2]point.Index
}
